package autogen

// AUTO GENERATED CODE DO NOT MODIFY MANUALLY
import (
	"github.com/gophercloud/gophercloud"

	"github.com/cloudwan/gohan/extension/gohanscript"
	"github.com/cloudwan/gohan/extension/gohanscript/lib"
)

func init() {

	gohanscript.RegisterStmtParser("get_openstack_client",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var authURL string
				iauthURL := stmt.Arg("auth_url", context)
				if iauthURL != nil {
					authURL = iauthURL.(string)
				}
				var userName string
				iuserName := stmt.Arg("user_name", context)
				if iuserName != nil {
					userName = iuserName.(string)
				}
				var password string
				ipassword := stmt.Arg("password", context)
				if ipassword != nil {
					password = ipassword.(string)
				}
				var domainName string
				idomainName := stmt.Arg("domain_name", context)
				if idomainName != nil {
					domainName = idomainName.(string)
				}
				var tenantName string
				itenantName := stmt.Arg("tenant_name", context)
				if itenantName != nil {
					tenantName = itenantName.(string)
				}
				var tenantID string
				itenantID := stmt.Arg("tenant_id", context)
				if itenantID != nil {
					tenantID = itenantID.(string)
				}
				var version string
				iversion := stmt.Arg("version", context)
				if iversion != nil {
					version = iversion.(string)
				}

				result1,
					err :=
					lib.GetOpenstackClient(
						authURL, userName, password, domainName, tenantName, tenantID, version)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("GetOpenstackClient",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			authURL, _ := args[0].(string)
			userName, _ := args[1].(string)
			password, _ := args[2].(string)
			domainName, _ := args[3].(string)
			tenantName, _ := args[4].(string)
			tenantID, _ := args[5].(string)
			version, _ := args[6].(string)

			result1,
				err :=
				lib.GetOpenstackClient(
					authURL, userName, password, domainName, tenantName, tenantID, version)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("openstack_token",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}

				result1 :=
					lib.OpenstackToken(
						client)

				return result1, nil

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackToken",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)

			result1 :=
				lib.OpenstackToken(
					client)
			return []interface{}{
				result1}

		})

	gohanscript.RegisterStmtParser("openstack_get",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}
				var url string
				iurl := stmt.Arg("url", context)
				if iurl != nil {
					url = iurl.(string)
				}

				result1,
					err :=
					lib.OpenstackGet(
						client, url)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackGet",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)
			url, _ := args[0].(string)

			result1,
				err :=
				lib.OpenstackGet(
					client, url)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("openstack_ensure",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}
				var url string
				iurl := stmt.Arg("url", context)
				if iurl != nil {
					url = iurl.(string)
				}
				var postURL string
				ipostURL := stmt.Arg("post_url", context)
				if ipostURL != nil {
					postURL = ipostURL.(string)
				}
				var data interface{}
				idata := stmt.Arg("data", context)
				if idata != nil {
					data = idata.(interface{})
				}

				result1,
					err :=
					lib.OpenstackEnsure(
						client, url, postURL, data)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackEnsure",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)
			url, _ := args[0].(string)
			postURL, _ := args[0].(string)
			data, _ := args[0].(interface{})

			result1,
				err :=
				lib.OpenstackEnsure(
					client, url, postURL, data)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("openstack_put",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}
				var url string
				iurl := stmt.Arg("url", context)
				if iurl != nil {
					url = iurl.(string)
				}
				var data interface{}
				idata := stmt.Arg("data", context)
				if idata != nil {
					data = idata.(interface{})
				}

				result1,
					err :=
					lib.OpenstackPut(
						client, url, data)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackPut",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)
			url, _ := args[0].(string)
			data, _ := args[0].(interface{})

			result1,
				err :=
				lib.OpenstackPut(
					client, url, data)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("openstack_post",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}
				var url string
				iurl := stmt.Arg("url", context)
				if iurl != nil {
					url = iurl.(string)
				}
				var data interface{}
				idata := stmt.Arg("data", context)
				if idata != nil {
					data = idata.(interface{})
				}

				result1,
					err :=
					lib.OpenstackPost(
						client, url, data)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackPost",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)
			url, _ := args[0].(string)
			data, _ := args[0].(interface{})

			result1,
				err :=
				lib.OpenstackPost(
					client, url, data)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("openstack_delete",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}
				var url string
				iurl := stmt.Arg("url", context)
				if iurl != nil {
					url = iurl.(string)
				}

				result1,
					err :=
					lib.OpenstackDelete(
						client, url)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackDelete",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)
			url, _ := args[0].(string)

			result1,
				err :=
				lib.OpenstackDelete(
					client, url)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("openstack_endpoint",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var client *gophercloud.ServiceClient
				iclient := stmt.Arg("client", context)
				if iclient != nil {
					client = iclient.(*gophercloud.ServiceClient)
				}
				var endpointType string
				iendpointType := stmt.Arg("endpoint_type", context)
				if iendpointType != nil {
					endpointType = iendpointType.(string)
				}
				var name string
				iname := stmt.Arg("name", context)
				if iname != nil {
					name = iname.(string)
				}
				var region string
				iregion := stmt.Arg("region", context)
				if iregion != nil {
					region = iregion.(string)
				}
				var availability string
				iavailability := stmt.Arg("availability", context)
				if iavailability != nil {
					availability = iavailability.(string)
				}

				result1,
					err :=
					lib.OpenstackEndpoint(
						client, endpointType, name, region, availability)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("OpenstackEndpoint",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			client, _ := args[0].(*gophercloud.ServiceClient)
			endpointType, _ := args[0].(string)
			name, _ := args[1].(string)
			region, _ := args[2].(string)
			availability, _ := args[3].(string)

			result1,
				err :=
				lib.OpenstackEndpoint(
					client, endpointType, name, region, availability)
			return []interface{}{
				result1,
				err}

		})

}
