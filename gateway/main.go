package main

import (
	"flag"
)

var addr = flag.String("addr", ":8080", "address to listen")
var authAddr = flag.String("auth_addr", "localhost:8081", "address for auth service")
var tripAddr = flag.String("trip_addr", "localhost:8082", "address for trip service")
var profileAddr = flag.String("profile_addr", "localhost:8082", "address for profile service")
var carAddr = flag.String("car_addr", "localhost:8084", "address for car service")

func main() {
	flag.Parse()

	//lg, err := server.NewZapLogger()
	//if err != nil {
	//	log.Fatalf("cannot create zap logger:%v", err)
	//}
	//c := context.Background()
	//c, cancel := context.WithCancel(c)
	//defer cancel()
	//
	//runtime.NewServeMux(runtime.WithMarshalerOption(
	//	runtime.MIMEWildcard, &runtime.JSONPb{
	//		EnumsAsInts: true,
	//		OrigName:    true,
	//	},
	//), runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
	//	if key == textproto.CanonicalMIMEHeaderKey(runtime.MetadataHeaderPrefix+auth.ImpersonateAccountHeader) {
	//		return "", false
	//	}
	//	return runtime.DefaultHeaderMatcher(key)
	//}))
	//
	//serverConfig := []struct {
	//	name         string
	//	addr         string
	//	registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opt []grpc.DialOption) (err error)
	//}{
	//	{
	//		name: "auth",
	//		addr: *authAddr,
	//	},
	//}

}
