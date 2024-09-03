/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package global

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitSysTips() {
	var logo = `
   ____   ____\_   _____/____    _________.__. /  _  \    __| _/_____ |__| ____  
  / ___\ /  _ \|    __)_\__  \  /  ___<   |  |/  /_\  \  / __ |/     \|  |/    \ 
 / /_/  >  <_> )        \/ __ \_\___ \ \___  /    |    \/ /_/ |  Y Y  \  |   |  \
 \___  / \____/_______  (____  /____  >/ ____\____|__  /\____ |__|_|  /__|___|  /
/_____/               \/     \/     \/ \/            \/      \/     \/        \/`
	var sys = `
Version: 1.0.0
Author: AnRuo
WebSite: https://www.kubesre.com/`
	var run = "Address: " + fmt.Sprintf("%s:%d", viper.GetString("server.address"),
		viper.GetInt("server.port"))
	fmt.Println(logo)
	fmt.Println(sys)
	fmt.Println(run)

}
