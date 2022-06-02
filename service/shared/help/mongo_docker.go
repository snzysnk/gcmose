package help

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"testing"
	"time"
)

func RunMongoDbInDocker(ApplyURI *string, m *testing.M) int {
	//创建docker客户端
	//client.FromEnv加载环境变量
	//client.WithAPIVersionNegotiation() 自动协商版本号
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	ctx := context.Background()
	if err != nil {
		panic(cli)
	}

	//ContainerCreate 有六个参数(ctx,config,hostConfig,networkingConfig,platform,containerName)
	//ctx 使用context.Background()即可
	//config 文件配置 重要的是docker端口曝露
	//hostConfig 本地文件配置 重要的是docker 和本地端口的映射
	//networkingConfig 网络设置
	//platform 平台设置
	//containerName 容器名称
	con, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "mongo:latest", //docker镜像名称
		ExposedPorts: nat.PortSet{ //曝露docker内部端口/连接方式
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{ //端口绑定设置
			"27017/tcp": []nat.PortBinding{ //将docker的27017以tcp方式绑定到本地的具体端口
				{
					HostIP:   "127.0.0.1", //127.0.0.1代表只有本机可以连接,0.0.0.0 所有机器可连
					HostPort: "0",         //0 代表docker自动寻找一个未被占用的本地端口连接 不是0就是设置固定值，如果端口占用就报错
				},
			},
		},
	}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	//ContainerStart 有三个参数
	//ctx 一般使用context.Background()即可
	//containerID 容器ID
	//options 启动参数
	err = cli.ContainerStart(ctx, con.ID, types.ContainerStartOptions{})

	if err != nil {
		panic(err)
	}

	fmt.Println("start")

	//ContainerInspect 两个参数
	//ctx 一般使用context.Background()即可
	//containerID 容器ID
	//代码写在 容器创建之后，停止之前
	inspect, err := cli.ContainerInspect(ctx, con.ID)

	if err != nil {
		panic(inspect)
	}

	//返回的是 nat.PortMap 其中 HostIP 是IP HostPort 是端口
	portInfo := inspect.NetworkSettings.Ports["27017/tcp"][0]

	fmt.Printf("启动mongodb在%s端口:%s", portInfo.HostIP, portInfo.HostPort)

	*ApplyURI = fmt.Sprintf("mongodb://%s:%s", portInfo.HostIP, portInfo.HostPort)

	time.Sleep(time.Second * 5)

	fmt.Println("结束mongodb")

	defer func() {
		timeOut := time.Second * 10
		//ContainerStop 有三个参数
		//ctx 一般使用context.Background()即可
		//containerID 容器ID
		//timeout 超时时间

		err = cli.ContainerStop(ctx, con.ID, &timeOut)

		if err != nil {
			panic(err)
		}

		//ContainerRemove 有三个参数
		//ctx 一般使用context.Background()即可
		//containerID 容器ID
		//options 配置参数
		err = cli.ContainerRemove(ctx, con.ID, types.ContainerRemoveOptions{})
		// 如果要跳过容器停止步骤,直接将容器移除,需要设置 options Force:true
		//err = cli.ContainerRemove(ctx, con.ID, types.ContainerRemoveOptions{Force: true})

		if err != nil {
			panic(err)
		}
	}()

	m.Run()

	return 0
}
