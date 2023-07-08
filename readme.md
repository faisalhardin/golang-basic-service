1. Run `make docker-up` in terminal 1. Please give it sometime to load.
2. Run `make consumer` in terminal 2. If there are no error log running on the terminal then it is safe to run the next step.
3. Run `make producer`  in terminal 3.
4. Run `make grpc-service` in terminal 4. Connect a client RPC service through port 9000. I use [grpcui](https://github.com/fullstorydev/grpcui) since my local machine run on M1 architecture. Else, can be tested with [grpcox](https://github.com/gusaul/grpcox)
