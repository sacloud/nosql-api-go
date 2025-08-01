# sacloud/nosql-api-go

Go言語向けのさくらのクラウド NoSQL APIライブラリ

NoSQL ドキュメント: https://manual.sakura.ad.jp/cloud/appliance/nosql/index.html

## 概要

sacloud/nosql-api-goはさくらのクラウド NoSQL APIをGo言語から利用するためのAPIライブラリです。

```go
package main

import (
	"context"
	"fmt"
	"net/netip"

	nosql "github.com/sacloud/nosql-api-go"
	v1 "github.com/sacloud/nosql-api-go/apis/v1"
)

func main() {
	client, err := nosql.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	databaseOp := nosql.NewDatabaseOp(client)
	// 以下はテスト用のセットアップ。アドレスなどは実際の環境向けに編集する
	resCreated, err := databaseOp.Create(ctx, v1.NosqlCreateRequestAppliance{
		Name:        "sdk-test-db",
		Description: v1.NewOptString("This is a test database"),
		Plan:        v1.NewOptPlan(v1.Plan{ID: v1.NewOptInt(51004)}),
		Tags:        v1.NewOptNilTags([]string{"nosql"}),
		Settings: v1.NosqlCreateRequestApplianceSettings{
			Backup: v1.NewOptNilNosqlCreateRequestApplianceSettingsBackup(v1.NosqlCreateRequestApplianceSettingsBackup{
				Connect:   "nfs://192.168.0.31/export",
				DayOfWeek: v1.NewOptNilNosqlCreateRequestApplianceSettingsBackupDayOfWeekItemArray([]v1.NosqlCreateRequestApplianceSettingsBackupDayOfWeekItem{"sun"}),
				Time:      v1.NewOptNilString("00:00"),
				Rotate:    v1.NewOptInt(4),
			}),
			SourceNetwork:    []string{},
			Password:         v1.Password("sdktest-12345"),
			ReserveIPAddress: netip.MustParseAddr("192.168.0.10"),
		},
		Remark: v1.NosqlRemark{
			// NosqlRemarkNosqlの設定は現状固定値なので、DefaultUser以外は変更しない
			Nosql: v1.NewOptNosqlRemarkNosql(v1.NosqlRemarkNosql{
				DatabaseEngine:  "Cassandra",
				DatabaseVersion: "4.1.9",
				DefaultUser:     "sdktest",
				DiskSize:        102400,
				Memory:          8192,
				Nodes:           3,
				Port:            9042,
				Storage:         "SSD",
				Virtualcore:     3,
				Zone:            "tk1b",
			}),
			Servers: []v1.NosqlRemarkServersItem{
				{UserIPAddress: netip.MustParseAddr("192.168.0.4")},
				{UserIPAddress: netip.MustParseAddr("192.168.0.5")},
				{UserIPAddress: netip.MustParseAddr("192.168.0.6")},
			},
		},
		UserInterfaces: []v1.NosqlCreateRequestApplianceUserInterfacesItem{
			{
				// 実際のSwitchのリソースIDを指定する
				Switch:         v1.NosqlCreateRequestApplianceUserInterfacesItemSwitch{ID: "111111111111"},
				UserIPAddress1: netip.MustParseAddr("192.168.0.4"),
				UserIPAddress2: v1.NewOptIPv4(netip.MustParseAddr("192.168.0.5")),
				UserIPAddress3: v1.NewOptIPv4(netip.MustParseAddr("192.168.0.6")),
				UserSubnet: v1.NewOptNosqlCreateRequestApplianceUserInterfacesItemUserSubnet(
					v1.NosqlCreateRequestApplianceUserInterfacesItemUserSubnet{
						DefaultRoute:   "192.168.0.1",
						NetworkMaskLen: 24,
					}),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Database: %+v\n", resCreated)

	res, err := databaseOp.Read(ctx, resCreated.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read Database: %+v\n", res)

	instanceOp := nosql.NewInstanceOp(client, res.ID)
	version, err := instanceOp.GetVersion(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Version: %+v\n", version)

	// Backup API
	// backupOp := nosql.NewBackupOp(client, res.ID)
}
```

NOTE: DatabaseAPIにあるNoSQL更新APIは設定のバリデーションのみで、実際に更新するには追加で反映APIを呼ぶ必要があります。 `Update` -> `ApplyChanges`.

:warning:  v1.0に達するまでは互換性のない形で変更される可能性がありますのでご注意ください。

## ogenによるコード生成

以下のコマンドを実行

```
$ go get -tool github.com/ogen-go/ogen/cmd/ogen@latest
$ go tool ogen -package v1 -target apis/v1 -clean -config ogen-config.yaml ./openapi/openapi.json
$ git apply fix-list-db-api.diff  # Listがogenの出力したコードではうまくいかないので修正
```

## License

`nosql-api-go` Copyright (C) 2025- The sacloud/nosql-api-go authors.
This project is published under [Apache 2.0 License](LICENSE).