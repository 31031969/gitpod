module github.com/gitpod-io/gitpod-local-cli

go 1.21

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gitpod-io/gitpod/gitpod-protocol v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/zalando/go-keyring v0.1.1
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/oauth2 v0.12.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2
	google.golang.org/grpc v1.58.2 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

require (
	github.com/bufbuild/connect-go v1.10.0
	github.com/dustin/go-humanize v1.0.1
	github.com/gitpod-io/gitpod/components/public-api/go v0.0.0-20231013185921-d2872f527de3
	github.com/go-git/go-git/v5 v5.9.0
	github.com/manifoldco/promptui v0.9.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.17.0
	github.com/toqueteos/webbrowser v1.2.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/danieljoos/wincred v1.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/godbus/dbus/v5 v5.0.3 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/sagikazarmark/locafero v0.3.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/skeema/knownhosts v1.2.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/sourcegraph/jsonrpc2 v0.0.0-20200429184054-15c2290dcb37 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gitpod-io/gitpod/gitpod-protocol => ../gitpod-protocol/go // leeway

replace github.com/gitpod-io/gitpod/local-app/api => ../local-app-api/go // leeway

replace github.com/gitpod-io/gitpod/supervisor/api => ../supervisor-api/go // leeway
