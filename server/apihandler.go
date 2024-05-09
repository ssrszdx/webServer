package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/tinode/chat/server/store"
	"github.com/tinode/chat/server/store/types"
	"net/http"
	"strconv"
	"strings"
)

// Response 是一个用于 JSON 响应的结构体
type Response struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    string `json:"data"`
}

type TResponse struct {
	Message string          `json:"message"`
	Code    string          `json:"code"`
	Data    []types.Coltest `json:"data"`
}

type SResponse struct {
	Message string          `json:"message"`
	Code    string          `json:"code"`
	Data    []types.UserExt `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 设置状态码
	w.WriteHeader(http.StatusOK)

	// 设置头部信息
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	// 发送文本响应
	// w.Write([]byte("Hello, World!"))

	// 发送 JSON 响应
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)

	// 发送 HTML 响应
	// w.Write([]byte("<html><body><h1>Hello, World!</h1></body></html>"))
}

// 登录
func syslogin(writer http.ResponseWriter, request *http.Request) {
	// 这里我们只是简单地检查用户名和密码是否匹配，实际应用中应该使用更安全的验证方式
	sflag := request.FormValue("sflag")
	if sflag == "1" {
		result := strings.Split(Mconfig.SysKey, "/")
		correctUsername := result[0] //"admin"
		correctPassword := result[1] //"pkuadmin"

		// 从请求中获取用户名和密码
		username := request.FormValue("username")
		password := request.FormValue("password")
		//response := Response{Message: "Hello, World!"}
		// 检查用户名和密码是否匹配
		if username == correctUsername && password == correctPassword {
			// 登录成功
			//fmt.Fprintln(writer, "登录成功，欢迎回来，管理员！")
			response := Response{Message: "登录成功！", Code: "1000", Data: ""}
			json.NewEncoder(writer).Encode(response)
		} else {
			// 登录失败
			response := Response{Message: "登录失败！", Code: "1001", Data: ""}
			json.NewEncoder(writer).Encode(response)
		}
	} else { //学生登录
		//var userOBJ types.UserExt
		testid, _ := strconv.ParseInt(request.FormValue("testid"), 10, 16)

		var userOBJ = types.UserExt{
			Username: request.FormValue("username"),
			Phone:    request.FormValue("phone"),
			Testid:   int16(testid),
			Clientid: request.FormValue("clientid"),
		}

		err := store.Userlogin(userOBJ)
		if err == nil {
			response := Response{Message: "登录成功！", Code: "1000", Data: ""}
			json.NewEncoder(writer).Encode(response)
		} else {
			response := Response{Message: "登录失败！", Code: "1001", Data: ""}
			json.NewEncoder(writer).Encode(response)
		}
	}

}

func createtest(writer http.ResponseWriter, request *http.Request) {
	// 从请求中获取测试名称和测试内容
	name := request.FormValue("name")
	parts := request.FormValue("parts")
	timestr := request.FormValue("timestr")
	// 创建测试
	recordid, err := store.Createtest(name, parts, timestr)
	if err == nil {
		response := Response{Message: "创建成功！", Code: "1000", Data: recordid}
		json.NewEncoder(writer).Encode(response)
	} else {
		response := Response{Message: "创建失败！", Code: "1001", Data: ""}
		json.NewEncoder(writer).Encode(response)
	}
}

func deletetest(writer http.ResponseWriter, request *http.Request) {
	// 从请求中获取测试名称和测试内容
	id := request.FormValue("id")
	var intId, _ = strconv.Atoi(id)
	// 创建测试
	err := store.Deletetest(intId)
	if err == nil {
		response := Response{Message: "删除成功！", Code: "1000", Data: ""}
		json.NewEncoder(writer).Encode(response)
	} else {
		response := Response{Message: "删除失败！", Code: "1001", Data: ""}
		json.NewEncoder(writer).Encode(response)
	}
}

// 获取测试列表
func gettestlist(writer http.ResponseWriter, request *http.Request) {
	//if request.Method == "GET" {
	sRds, err := store.GetColTest()
	jsonSRDS, _ := json.Marshal(sRds)
	if err == nil {
		fmt.Println("store.GetColTest() :", string(jsonSRDS))
		response := TResponse{Message: "调用成功！", Code: "1000", Data: sRds}
		json.NewEncoder(writer).Encode(response)
	} else {
		response := Response{Message: "调用失败！", Code: "1001", Data: string(jsonSRDS)}
		json.NewEncoder(writer).Encode(response)
	}
	//}
}

// 更新列表
func updatetest(writer http.ResponseWriter, request *http.Request) {
	// 从请求中获取测试名称和测试内容
	id := request.FormValue("id")
	name := request.FormValue("name")
	parts := request.FormValue("parts")
	timestr := request.FormValue("timestr")
	var intId, _ = strconv.Atoi(id)

	var coltest = types.Coltest{
		Id:      int16(intId),
		Name:    name,
		Parts:   parts,
		Timestr: timestr,
	}

	// 创建测试
	err := store.Updatetest(coltest)
	if err == nil {
		response := Response{Message: "修改成功！", Code: "1000", Data: ""}
		json.NewEncoder(writer).Encode(response)
	} else {
		response := Response{Message: "修改失败！", Code: "1001", Data: ""}
		json.NewEncoder(writer).Encode(response)
	}
}

// getColInfo 获取分组信息
// 获取测试列表
func getgroupinfo(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		clientid := request.FormValue("clientid")
		testid := request.FormValue("testid")
		sRds, err := store.GetGroupInfo(clientid, testid)
		jsonSRDS, _ := json.Marshal(sRds)
		if err == nil {
			fmt.Println("store.GetColTest() :", string(jsonSRDS))
			response := SResponse{Message: "调用成功！", Code: "1000", Data: sRds}
			json.NewEncoder(writer).Encode(response)
		} else {
			response := Response{Message: "调用失败！", Code: "1001", Data: string(jsonSRDS)}
			json.NewEncoder(writer).Encode(response)
		}
	}
}

// 直接执行sql
func dbExec(dbSql string) {
	dsn := "root:@tcp(127.0.0.1:3306)/tinode"

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接测试失败:", err)
		return
	}
	// 查询表
	rows, err := db.Query(dbSql)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	defer rows.Close()

	type record struct {
		id    int
		name  string
		parts string
	}

	// 遍历结果
	var tests []record
	for rows.Next() {
		var test record
		err = rows.Scan(&test.id, &test.name, &test.parts)
		if err != nil {
			fmt.Println("读取数据失败:", err)
			return
		}
		tests = append(tests, test)
	}

	// 打印结果
	for _, test := range tests {
		fmt.Printf("ID: %d, Name: %s, Parts: %s\n", test.id, test.name, test.parts)
	}

}
