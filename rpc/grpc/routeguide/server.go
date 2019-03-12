
package main

import (
    "context"
    srv "com/beyanger/service"
)

type RouteGuideServer struct {

}

func (s *RouteGuideServer) GetFeature(ctx context.Context, point *srv.Point) (f *srv.Feature, err error) {
    return
}

func (s *RouteGuideServer) ListFeatures(rect *srv.Rectangle, stream srv.RouteGuide_ListFeaturesServer) error {
    return nil
}

func (s *RouteGuideServer) RecordRoute(stream srv.RouteGuide_RecordRouteServer) error {
    return nil
}

func (s *RouteGuideServer) RouteChat(stream srv.RouteGuide_RouteChatServer) error {
    return nil
}

func main() {

}

