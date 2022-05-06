package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
environment: D
app: booking-identity-management
grpc_address: 20000
http_address: 20001
sensitive_fields:
sensitive_fields_str: password:MASKALL
legacy_token: true
auth:
  authorizations:
    enabled: false
mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: 
  database: giang
  is_migrate: true
paseto: 
  refresh_token: false
  expired_time: 21600
  issuer: BIM
  public_key: 7373682D65643235353139204141414143334E7A6143316C5A4449314E544535414141414949736C6A427337457A337A38447666537A592F7954774F61653939725567654E4742472B352F6C564A3679206769616E676C653231393840676D61696C2E636F6D
  private_key: 2D2D2D2D2D424547494E204F50454E5353482050524956415445204B45592D2D2D2D2D0D0A6233426C626E4E7A614331725A586B74646A45414141414142473576626D554141414145626D39755A5141414141414141414142414141414D7741414141747A633267745A570D0A51794E5455784F5141414143434C4A5977624F784D39382F41373330733250386B38446D6E7666613149486A5267527675663556536573674141414A67527152625045616B570D0A7A7741414141747A633267745A5751794E5455784F5141414143434C4A5977624F784D39382F41373330733250386B38446D6E7666613149486A5267527675663556536573670D0A41414145435352616B4559335631445A4A5868737456744F785461656636664148524550644471506270496E76494649736C6A427337457A337A38447666537A592F7954774F0D0A61653939725567654E4742472B352F6C564A367941414141465764705957356E624755794D546B345147647459576C734C6D4E7662513D3D0D0A2D2D2D2D2D454E44204F50454E5353482050524956415445204B45592D2D2D2D2D0D0A
jwt:
  private_key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS2dJQkFBS0NBZ0VBMnc2SDRCTGttU2tzNFhhV1h1bDFNSndqdjhpTTUzU2dNdlNnMktZYnlhSEhuUTg5Ck5HQ2poVVBFQkJDcHkvNEVQMkxCVFgvK2xxUWdsS3U4em1NbGx3YWNzWnB0NWkrdk94Q1NFV2JvQThrQVZndHEKSlp5eVl5QkJBZDhINGVHZzlRS0RsS3hLVE1udFE0bFRmT0gyOHRRVXcrYnFVc3NWblFwR3g2TlM4bVpxamQrSQp5YW04ekRpWmdyY3IvMmo3RTFrWXdFb0VvUkxkUUhyb3lmRlNWSEtKeWVVZjVNUFFSWVlEbGdEc0pqb0o3RXFtCkdTZ2dvM01lM2JVM2M5RGRQdkYyVnUrR0lyZDFCU3ZWaURXd1p4MnYvNlAxc0lNQXYrT2xkbVVEaVRsVStNanMKK2ZEK0FBY01uYXdtRzJ2UUtqaHM1RlNTZ1hXUFR3KzRpNDd2Y1N0V0lRaWlyZmdUd1ZtOXNCWHdiYjN1Sy9kVwozSi9MNXhocWIzYWs3ZGtvZXFpd2poeWNhMGdNRy9SVW5HMlFLTzl3dURFeFE5SVFzSGJiSXpwdElqQlVidmJFCk5md2l0ekNweE9hRjYzOUJCclJidHhHMjFpT3FJa1FUVEc4Z2huKzRORG56MThRUGFsZEdjM2ZPaERpYjZjaXIKK2FyUEJpK1FsWDZIQmhzK2ZIb251Q20rMjMyNnZkMmpIYmowem1jY3hsM3JuMXVFYVJINlgvOHZzbGZYejhySQpBWTlpdUl3aGIwYjBaYW93SXE0cFdYM1dUZUd2cTFBNEtnMXFOdm9kbTM5OWZJS25YbWVQc3RZaWJXMm5IK3hMCnpsemRNN1ZtWFB6cjA3L1BvQzhnVkVySERDeVRlajdIY2s0SUZhc0U2aXMzRTRQTDc5U3p1M2R5NVlrQ0F3RUEKQVFLQ0FnRUF4UTlSVkZZcXAwM1BQaFE4T2EyeTZzUEx0UjlwQWw1OEx0cTZSNnRHbkNDWlBjVVR2aWFoYzZUKwpsTm1ZcGhZaEpCUm02SjBZSGRRUUt5NURiemJwL2U3SXoweXBTTVhrMFRHYTlXU3BtNTRhVzAxeGM5eCs3WVY2CmRZQjV5czM2eWU0MUNHM2xlU2E5OWw5M1o0U3Zic3N5MHZqTUd5YnVqcE5NRVJQdjduZ0xxdzFJTklJcGdWMzUKV3dvRTNvTkdPcVdRQ291TGxvc3QxbSs3TXMrUXpMOHREUjY5amtKSS9nSU0vOGhjcUJiV0RtY28vTUNJZkFGMwo1dUs2dThsTS91NmhFcHZ2QXhRZWhVUmxsMThWRGU0N0FiSXh0MVJpUEQxQ0FEdm03bTlRN1cxeklYV1dTZGdOCmxWbmovZzhDK0xoQjNUWG5oZWFId1F2WGJGUnRPS3JuYmVXZ0NjRzhad2F5L1l5Sm5pSkdOODZ4dU55YUdKaGQKVWh1REJ6NHRuRkNYY1JwNTh1UURjL1JQMlVmNUNEbndmOFNMTi9hTG1nSGZROEU4bVZ1R2Jmc3Y2TXV6RjcxbQpMNFluMW9wbFNpbit0WVhnaHdKZkd6WTREWElzZy9aYXh1U0ZQVDJVT2xCVXZ4cnRTdTdTQk9kQ0kyQTFRWnppClNFOFYrSmJYMHhOYzBWYjJ1cmt5T0hPak5xUkZYY2ZsZUx4eW9tRU5xZE5xUmFwS0szeEd4YWZRK1prOEVjUk0KMFAyQkVnQzJ4TGNYL0JkTjVoMEJISFRFd0JkeEgxU0xpYkEydFJVZEJqMlVIVmVVQlpVb091bDhNc1hVYlBzOAo2S25pYno2M3hNMUxjTTErOEhJSGF3elpoZEF0K1JRcktoK0RQWSt6TlZmN0NISmlDQUVDZ2dFQkFQWjZ2OUdoCk11VEtFRHFJL0FnQUJ6b2JNc2hsaEFaWkVlT0JmZjd6N3RncjRtVFZBYWVNR0RwbVM3SnFyOFVPN3N1QmdOSkEKOE82WVRGWFoxeXhoMWlCTnR4bGRadUdhRGJGb1Z2aEI0UW4vSCs0R1d0d1FqNUt0RDBFR3Axd3NUTjBSSEVUbQoyZnJHQmIyazAvaGhHdm1uK3dXdGc4bVNsVWlDNkw2Q2IvWmJhSlFpblZIL05NTkdyUmNRc3hvcnJnM29KOTJmCndjbDJuTGFDVDFNeUJZdnRkSVY4clppUzgzL0hycDhGcVVJZ2xWaVpFaW4yVTU4dWNrZVBHbTNZdHo1TTlyN1YKeHZBRHd6R2NoajZDSi9EV01rMmFKOGFuNHY0UERjSUFweVBPckM5d3JPNE5lc25kTFNnQzN3Uk5Bd2Y0d1dicAo3SURyTHA4WVoxYUFJc2tDZ2dFQkFPT0VubUl0S291bEZXbnJCZ0MrQ1ZUU2g3VmZETVl0ajM2aDdiSGp2R3gyCkx2dVJ3aDVxSFpuSmZVVWNTWXc4VG5kbitZaDQ0REd6bmR0UXlaUGVycndWVmhhZnRhUiswUDNtc0xyZnYvaDIKZFJCU00zNk9JUWwzbzcyM1M2cXVHUUVPM2dTR25qWGUzcnhXcE4yOUc5MzErNytCeXVQczc0d1RYZHJLZEVuaApoZ2h5RG92cWNONGtjSHpPeCtTMmRwakwyT2VGdWZMRGY4RitVQy82VUtsT2pKcTgzTHdITGtqakp2ZE1wdE9iCkJKMmJUN2hMRFBRV2JycnlYeEZZdmlaeWtHRkRjMCt3dElwbmtWbThHOVQwa2FCNExReEFMUTh2Ums3SzVlRFgKUlJib1JHaFloUDlUOUMwcjF0UUYwckZYVGZFa1ZMSXR3NHBVbmpmZFRNRUNnZ0VCQUxZeHJqRUs4MC9qZlhPSApacVdUUDhSWDdvWHJaRmc4NXRoQkZCMnR1dFBZUEpWRVR1REcwZkg1UStNeXNaTDhvcytUZnVIc1Z5eEhMOFlXCkxOUUZ6ZjQ2OUNtdnZTanVjMGk4VSsxVzVUNDZkQ1BDVnM5cy9uUHhzT0RrUmpvQUZTRWtVTDB0MTl0akhVTTkKMytJdUJYNmVDUlA5V3ZsWmZUZ0liSHJHZUhTbXhQc2JTeWNkZXFjT1MweE5iOWJjaXVDRXgxTXZrRk9YZGVnRApFeGZ3VmVGRk9XaWk3TDluZ003bmo5REVvb1RiQ1FwZ3UraUF5UGRWb2V2N2dCSEtyRWxGSXUrVm5tUTBvOGk4Clp3ODFyc2g3ek50SUpEYllTeFo1a2pYaHZhcmVUckVTblFHcXZTVVNkUnFxWktnTk5UeGpTTE1HSDJoOGNBM3kKUUJEVXp1RUNnZ0VCQUxCSXdsdDVlWHluMVpheTFuemV3ZVlxRmFKdldqdDNjN1ovaXVlSmVDWmx6L04yTlgxUgpaVXdTTlhER0FTcWFVNFhyVnducDluUXNJc1NhclFWWXhmVFBaMXV5TlpDVi9pS3Npd1RRZHpMeHZiUitySTJEClhJKzBWbmVadlJRclB3NTEwWXFhTDUyMzZNYnZMMkRydE5yZlRFZFBMc0E1aGtYM2Y0SVYyZEtnM2pSUEEzeGYKazBKYjQ1aWdzM3MvOFZaa285U0g1ZEwyclZka0lvZDZJT04rREMyYzVqNytKVnV1UFRQaC9XVDV1TExSY3d4dApna1RoMkx3YkxaUUlPWlVhc1VuYmRFNHJzVWVJQmJRTXIxZnFnNmRkVmpJbXk0TDNjemQ3dUcvay9CQkt1UzVNCllBaGhxVkVTcGFxME5GTGFQWGI1V25LUVNQRTJyVHpGeUVFQ2dnRUFFSSs4Uk9mbUVLck0wMWJGVUE3NUVxMVMKOW42bFhiczkrWkFIblNEK2t6M1RsYUlUcnpVd25mbTNKTkkrZENWRCs2YnNvWnpzcGFtSlNKVDEzdk5ja2hBVApCNmlmaUs1blgxb3RCS055aDA0Q3NFNmM5WkVRb0xVbWUrVXlxSlJla1UvZzg4VG1KNkdONnFKRUh1d0ZYMG1HCkZIYlVjNkVKOXY3OHdkNXgzVEtMcXNwNnkwTFhabVV5NUZpQWN2SVRHdnBVOUgxS2tITTZtWnZCZjdkSlBOd28KSnA4c1FXQ3NSMlB2ZjB0aUltU0puRmFXV1JnVXVsSE5KdU12R29CN2dJTVF1YXRicmtoU3hHb0Z3WXRuUk5YbAovNytZVkFhQ1VUdlVOUm1MREZPTkJaRW5lMitVclorVWJWVXQzSmJKenNoS25mNnJiaGVEOHl0OVpvZkpQdz09Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t
  public_key: LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUEydzZINEJMa21Ta3M0WGFXWHVsMQpNSndqdjhpTTUzU2dNdlNnMktZYnlhSEhuUTg5TkdDamhVUEVCQkNweS80RVAyTEJUWC8rbHFRZ2xLdTh6bU1sCmx3YWNzWnB0NWkrdk94Q1NFV2JvQThrQVZndHFKWnl5WXlCQkFkOEg0ZUdnOVFLRGxLeEtUTW50UTRsVGZPSDIKOHRRVXcrYnFVc3NWblFwR3g2TlM4bVpxamQrSXlhbTh6RGlaZ3Jjci8yajdFMWtZd0VvRW9STGRRSHJveWZGUwpWSEtKeWVVZjVNUFFSWVlEbGdEc0pqb0o3RXFtR1NnZ28zTWUzYlUzYzlEZFB2RjJWdStHSXJkMUJTdlZpRFd3Clp4MnYvNlAxc0lNQXYrT2xkbVVEaVRsVStNanMrZkQrQUFjTW5hd21HMnZRS2poczVGU1NnWFdQVHcrNGk0N3YKY1N0V0lRaWlyZmdUd1ZtOXNCWHdiYjN1Sy9kVzNKL0w1eGhxYjNhazdka29lcWl3amh5Y2EwZ01HL1JVbkcyUQpLTzl3dURFeFE5SVFzSGJiSXpwdElqQlVidmJFTmZ3aXR6Q3B4T2FGNjM5QkJyUmJ0eEcyMWlPcUlrUVRURzhnCmhuKzRORG56MThRUGFsZEdjM2ZPaERpYjZjaXIrYXJQQmkrUWxYNkhCaHMrZkhvbnVDbSsyMzI2dmQyakhiajAKem1jY3hsM3JuMXVFYVJINlgvOHZzbGZYejhySUFZOWl1SXdoYjBiMFphb3dJcTRwV1gzV1RlR3ZxMUE0S2cxcQpOdm9kbTM5OWZJS25YbWVQc3RZaWJXMm5IK3hMemx6ZE03Vm1YUHpyMDcvUG9DOGdWRXJIREN5VGVqN0hjazRJCkZhc0U2aXMzRTRQTDc5U3p1M2R5NVlrQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==
  is_usage_refresh_token: false
  refresh_token_expire: 21600
  short_live_token_expire: 180
  long_live_token_expire: 21600
  signing_method: RS512
  issuer: BIM
excluded_paths:
  - DoLogin
  - DoRegister
  - DoLogout
  - DoVerifyJWTToken
cors_local: true
`)

type (
	// Config represents config
	Config struct {
		Base `mapstructure:",squash"`
	}

	Base struct {
		Environment        string            `mapstructure:"enviroment"`
		App                string            `mapstructure:"app"`
		GRPCAddress        int               `mapstructure:"grpc_address"`
		HTTPAddress        int               `mapstructure:"http_address"`
		SensitiveFieldsStr string            `mapstructure:"sensitive_fields_str" replace_for:"sensitive_fields" config_type:"map"`
		SensitiveFields    map[string]string `mapstructure:"sensitive_fields"`
		LegacyToken        string            `mapstructure:"legacy_token"`
		Auth               auth              `mapstructure:"auth"`
		MySQL              mysql             `mapstructure:"mysql"`
		Paseto             paseto            `mapstructrue:"paseto"`
		JWT                jwt               `mapstructure:"jwt"`
		ExcludedPaths      []string          `mapstructure:"excluded_paths"`
		CORSLocal          bool              `mapstructure:"cors_local"`
	}

	auth struct {
		Authorizations string `mapstructure:"authorization"`
	}

	mysql struct {
		Host      string `mapstructure:"host"`
		Port      int    `mapstructure:"port"`
		Username  string `mapstructure:"username"`
		Password  string `mapstructure:"password"`
		Database  string `mapstructure:"database"`
		IsMigrate bool   `mapstructure:"is_migrate"`
	}

	paseto struct {
		PublicKey    string `mapstructure:"public_key"`
		PrivateKey   string `mapstructure:"private_key"`
		Issuer       string `mapstructure:"issuer"`
		ExpriredTime uint32 `mapstructure:"expired_time"`
		RefreshTOken string `mapstructure:"refresh_token"`
	}

	jwt struct {
		RefreshTokenExpireTime   uint32 `yaml:"refresh_token_expire" mapstructure:"refresh_token_expire"`
		ShortLongTokenExpireTime uint32 `yaml:"short_live_token_expire" mapstructure:"short_live_token_expire"`
		LongLiveTokenExpireTime  uint32 `yaml:"long_live_token_expire" mapstructure:"long_live_token_expire"`
		IsUsageRefreshToken      bool   `yaml:"is_usage_refresh_token" mapstructure:"is_usage_refresh_token"`
		Issuer                   string `yaml:"issuer" mapstructure:"issuer"`
		SigningMethod            string `yaml:"signing_method" mapstructure:"signing_method"`
		PrivateKey               string `yaml:"private_key" mapstructure:"private_key"`
		PublicKey                string `yaml:"public_key" mapstructure:"public_key"`
	}
)

func Load() (*Config, error) {
	var cfg = &Config{}
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
