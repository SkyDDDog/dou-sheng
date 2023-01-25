package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Version string `json:"version"` // 版本
	Weight  int64  `json:"weight"`  // 权重
}

func BuildPrefix(server Server) string {
	if server.Version == "" {
		return fmt.Sprintf("/%s/", server.Name)
	}
	return fmt.Sprintf("/%s/%s/", server.Name, server.Version)
}

func BuildRegisterPath(server Server) string {
	return fmt.Sprintf("%s%s", BuildPrefix(server), server.Addr)
}

// ParseValue	将value值反序列化到Server实例中
func ParseValue(value []byte) (Server, error) {
	server := Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, err
	}
	return server, nil
}

func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("Path Invalid")
	}
	server.Addr = strs[len(strs)-1]
	return server, nil
}

func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}
