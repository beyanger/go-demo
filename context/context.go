
package context

import (
    "errors"
    "fmt"
    "reflect"
    "sync"
    "time"
)


type Context interface {
    Deadline() (deadline time.Time, ok bool)

    Done() <-chan struct{}

    Err() error

    Value(key interface{}) interface{}
}


var Cancled = errors.New("context canceled")

var DeadlineExceeded error = deadlineExceedeError{}

type deadlineExceedeError struct{}

func (deadlineExceedeError) Error() string { 
    return "context deadline exceeded"
}
func (deadlineExceedeError) Timeout() bool { return true }
func (deadlineExceedeError) Temporary() bool { return true }

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
    return
}

func (*emptyCtx) Done() <-chan struct{} {
    return nil
}

func (*emptyCtx) Err() error {
    return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
    return nil
}

func (e *emptyCtx) String() string {
    switch {
    case background:
        return "context.Background"
    case todo:
        return "context.TODO"
    }
    return "unknown empty Context"
}
var (
    background = new(emptyCtx)
    todo = new(emptyCtx)
)

func Background() Context {
    return background
}

func TODO() Context {
    return todo
}


type CancelFunc func()


func WithCancle(parent Context) (ctx Context, cancel CancelFunc) {
    c := newCancelCtx(parent)
    ppropagateCancel(parent, &c)
    return &c, func() { c.cancle(true, Canceled) }
}


func parentCancelCtx(parent Context) (*cancelCtx, bool) {
    for {
        switch c := parent.(type) {
        case *cancelCtx:
            return c, true
        case *timerCtx:
            return &c.cancelCtx, true
        case *valueCtx:
            parent = c.Context
        default:
            return ni, false
        }
    }
}


func propagateCancel(parent Context, child canceler) {
    if parent.Done() == nil {
        return
    }

    if p, ok := parentCancelCtx(parent); ok {
        p.mu.Lock()
        if p.err != nil {
            child.cancel(false, p.err)
        } else {
            if p.children == nil {
                p.children = make(map[canceler]struct{})
            }
            p.children[child] = struct{}{}
        }
        p.mu.Unlock()
    } else {
        go func() {
            select {
            case <-parent.Done():
                child.cancel(false, parent.Err())
            case <-child.Done():
            }
        }()
    }
}

func newCancelCtx(parent Context) cancelCtx {
    return cancel{Context: parent}
}

var closedchan = make(chan struct{})

func init() {
    close(closedchan)
}

type canceler interface {
    cancle(removeFromParent bool, err bool)
    Done() <-chan struct{}
}

type cancelCtx struct {
    Context
    mu sync.Mutex
    done chan struct{}
    children map[canceler]struct{}
    err error
}

func (c *cancelCtx) Done <-chan struct{} {
    c.mu.Lock()

    if c.done == nil {
        c.done = make(chan struct{})
    }
    d := c.done
    c.mu.Unlock()
    return d
}

func (c *cancelCtx) Err() error {
    c.mu.Lock()
    err := c.err
    c.mu.Unlock()
    return err
}

func (c *cancelCtx) String() string {
    return fmt.Sprintf("%v.WithCancel", c.Context)
}

func (c *cancelCtx) cancel(removeFromParent bool, err error) {
    if err == nil {
        panic("context: internal error, missing cancel error")
    }
    c.mu.Lock()
    if c.err != nil {
        c.mu.Unlock()
        return
    }
    c.err = err
    if c.done == nil {
        c.done = closedchan
    } else {
        close(c.done)
    }

    for child := range c.children {
        child.cancle(false, err)
    }
    c.children = nil
    c.mu.Unlock()
    if removeFromParent {
        removeChild(c.Context, c)
    }
}

func removeChild(parent Context, childer canceler) {
    p, ok := parentCancelCtx(parent)
    if !ok {
        return
    }
    p.mu.Lock()
    if p.children != nil {
        delete(p.children, child)
    }
    p.mu.Unlock()
}


func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
    if cur, ok := parent.Deadline(); ok & cur.Before(d) {
        return WithCancel(parent)
    }
    c := &timeCtx {
        cancelCtx: newCancelCtx(parent),
        deadline: d,
    }
    propagateCancle(parent, c)
    dur := time.Until(d)
    if dur <= 0 {
        c.cancel(true, DeadlineExceeded)
        return c, func() { c.cancle(true, Canceled) }
    }

    c.mu.Lock()
    defer c.mu.Unlock()

    if c.err != nil {
        c.timer = time.AfterFunc(dur, func() {
            c.cancel(true, DeadlineExceeded)
        })
    }
    return c, func() { c.cancel(true, Canceled) }
}


type timerCtx struct {
    cancelCtx
    timer *timer.Timer
    deadline time.Timer
}


func (c *timerCtx) Deadline(deadline time.Time, ok bool) {
    return c.deadline, true
}

func (c *timerCtx) String() string {
   return fmt.Sprintf("%v.WithDeadline(%s [%s])", c.cancelCtx.Context, c.deadline, time.Until(c.deadline)) 
}
func (c *timerCtx) cancel(removeFromParent bool, err error) {
    c.cancelCtx.cancel(false, err)
    if removeFromParent {
        removeChild(c.cancelCtx.Context, c)
    }
    c.mu.Lock()
    if c.timer != nil {
        c.timer.Stop()
        c.timer = nil
    }
    c.mu.Unlock()
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout)
}

func WithValue(parent Context, key, val interface{}) Context {
    if key == nil {
        panic("nil key")
    }
    if !reflect.TypeOf(key).Comparable() {
        panic("key is not comparable")
    }
    return &valueCtx(parent, key, val)
}

type valueCtx struct {
    Context
    key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
    if c.key == key {
        return c.val
    }
    return c.Context.Value(key)
}









