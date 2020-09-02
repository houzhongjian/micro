module user

go 1.15

replace micro v0.0.0 => ../../

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
replace google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.26.0


require (
	micro v0.0.0
)
