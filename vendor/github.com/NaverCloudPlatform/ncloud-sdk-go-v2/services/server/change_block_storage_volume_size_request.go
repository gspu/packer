/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type ChangeBlockStorageVolumeSizeRequest struct {

	// 블록스토리지인스턴스번호
BlockStorageInstanceNo *string `json:"blockStorageInstanceNo"`

	// 블록스토리지사이즈
BlockStorageSize *int64 `json:"blockStorageSize"`
}