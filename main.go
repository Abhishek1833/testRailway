package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//"zocketAssignment/controllers"

	controllers "github.com/Abhishek1833/ItemList/Controllers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func init() {
	os.Setenv("PORT", "8000")
	setUpViper()
	registerDatabase()
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/itempost", controllers.PostItems).Methods("POST")
	r.HandleFunc("/item", controllers.GetItems).Methods("GET")
	r.HandleFunc("/item", controllers.UpdateItem).Methods("PUT")
	r.HandleFunc("/item", controllers.DeleteItem).Methods("DELETE")
	r.HandleFunc("/", controllers.Welcome)
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, r))
}
func registerDatabase() {
	runmode := cast.ToString(viper.Get("runmode"))
	mysql := viper.Get(runmode + ".mysql").(map[string]interface{})
	mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string)
	//mysqlConf := "mysql://root:5ku0BhbZ1VxhyxY2kJ8G@containers-us-west-74.railway.app:5542/railway"
	log.Println("conf", mysqlConf)
	orm.RegisterDataBase("default", "mysql", mysqlConf)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Kolkata")
	orm.Debug = true

}

//set up config file from conf folder
func setUpViper() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("global")
}
