package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"git.taservs.net/platform/edca-api/internal/api/controller"
	"git.taservs.net/platform/edca-api/internal/constant"
	edcaerror "git.taservs.net/platform/edca-api/internal/error"
	"git.taservs.net/platform/edca-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

const roleDataMD = `
| Endpoint                                                                                                    | EDCAUserManager | CustomerSupport | InstanceManager | FieldSupport | ContractSupport | FirmwareManager | ProductManager | SuperAdmin | DeviceManager | MarketingUser | AgencyViewer |
|-------------------------------------------------------------------------------------------------------------|-----------------|-----------------|-----------------|--------------|-----------------|-----------------|----------------|------------|---------------|---------------|--------------|
| GET,/api/agencies                                                                                           |                 | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          |               | TRUE         |
| POST,/api/agencies                                                                                          |                 |                 | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId                                                                                 |                 | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          |               | TRUE         |
| PATCH,/api/agencies/:agencyId                                                                               |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/device-stats                                                                    |                 | TRUE            | TRUE            |              | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/agencies/features                                                                                  |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/features                                                                        |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/features/:featureId                                                            |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/features/:featureId                                                          |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/storage                                                                         |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| PATCH,/api/agencies/:agencyId/storage/bytes                                                                 |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/storage/device                                                                 |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/storage/device                                                               |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/license-counts                                                                  |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| PATCH,/api/agencies/:agencyId/license-counts                                                                |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/lte-configs                                                                     |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/lte-configs                                                                    |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/auth-clients                                                                    |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/auth-clients/:authClientId                                                   |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/licenses                                                                        |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/licenses/:licenseId                                                          |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/licenses                                                                       |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/return-tickets                                                                  |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/return-tickets/users/:userId/tickets/:ticketId                                  |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/return-tickets/users/:userId/tickets/:ticketId/complete                        |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/return-tickets/users/:userId/tickets/:ticketId/resubmit                        |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/returns-tickets/users/:userId/tickets/:ticketId/tracking                       |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/returns-tickets/users/:userId/tickets/:ticketId/cancel                         |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/returns-tickets/users/:userId/tickets/:ticketId/items/:itemId                |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/returns-tickets/users/:userId/tickets/:ticketId/non-serialized-items/:itemId |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/users/:userId                                                                   | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/agencies/:agencyId/users/:userId/licenses                                                          | TRUE            | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/users/:userId/licenses/:licenseId                                            | TRUE            | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/audit-trail                                                                    | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/agencies/:agencyId/roles                                                                           | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/agencies/:agencyId/users/:userId/mfa-status                                                        | TRUE            | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/mfa-policy                                                                      | TRUE            | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/devices/:deviceId                                                               |                 | TRUE            | TRUE            |              | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| POST,/api/agencies/:agencyId/devices/:deviceId/status/deregistering                                         |                 | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       | TRUE          |               |              |
| PATCH,/api/agencies/:agencyId/users/:userId                                                                 | TRUE            | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/ranks                                                                           | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/agencies/:agencyId/drone-sense-integration                                                         |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| PATCH,/api/agencies/:agencyId/drone-sense-integration                                                       |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/users/:userId/reset-mfa                                                        | TRUE            | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/users/:userId/unlock                                                           | TRUE            | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/users/bulk-set-status                                                          | TRUE            | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/add-admin                                                                      | TRUE            | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/status                                                                         |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/change-account-type                                                            |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/cad-rms-integration/config                                                      |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| PATCH,/api/agencies/:agencyId/cad-rms-integration/config                                                    |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/cad-rms-integration/categories                                                 |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/cad-rms-integration/categories                                                  |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/cad-rms-integration/categories/template                                         |                 | TRUE            | TRUE            | TRUE         |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/feature-layers/:layerId/files/:layerFileId                                      |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/users                                                                                              | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| POST,/api/users                                                                                             | TRUE            |                 |                 |              |                 |                 |                | TRUE       |               |               |              |
| POST,/api/users/reinvite                                                                                    | TRUE            | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       |               |               |              |
| POST,/api/users/reset-password                                                                              | TRUE            | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| POST,/api/users/:userId/audit-trail                                                                         | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/devices                                                                                            |                 | TRUE            | TRUE            |              | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| POST,/api/devices/bulk-transfer                                                                             |                 | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       | TRUE          |               |              |
| POST,/api/devices/assign                                                                                    |                 | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       | TRUE          |               |              |
| POST,/api/devices/send-command                                                                              |                 |                 |                 |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/devices/logs                                                                                       |                 | TRUE            | TRUE            |              |                 |                 |                | TRUE       | TRUE          |               |              |
| GET,/api/devices/logs/download                                                                              |                 | TRUE            | TRUE            |              |                 |                 |                | TRUE       | TRUE          |               |              |
| POST,/api/devices/password/decrypt                                                                          |                 | TRUE            | TRUE            |              | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| GET,/api/devices/commands/model/:model                                                                      |                 |                 |                 |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/announcements                                                                                      |                 |                 |                 |              |                 |                 |                | TRUE       |               | TRUE          |              |
| GET,/api/announcements/:announcementId/thumbnail                                                            |                 |                 |                 |              |                 |                 |                | TRUE       |               | TRUE          |              |
| GET,/api/announcements/active/metadata                                                                      |                 |                 |                 |              |                 |                 |                | TRUE       |               | TRUE          |              |
| GET,/api/announcements/:announcementId                                                                      |                 |                 |                 |              |                 |                 |                | TRUE       |               | TRUE          |              |
| POST,/api/announcements                                                                                     |                 |                 |                 |              |                 |                 |                | TRUE       |               | TRUE          |              |
| PUT,/api/announcements/:announcementId                                                                      |                 |                 |                 |              |                 |                 |                | TRUE       |               | TRUE          |              |
| GET,/api/cad-rms-integration/cad-jobs                                                                       |                 | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/cad-rms-integration/cad-jobs/:id/download                                                          |                 | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/profiles                                                                                           | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| PUT,/api/profiles/update-password                                                                           | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| GET,/api/profiles/mfa/totp                                                                                  | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| POST,/api/profiles/mfa/issue                                                                                | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| POST,/api/profiles/mfa/verify                                                                               | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| DELETE,/api/profiles/mfa                                                                                    | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| POST,/api/profiles/otp/send                                                                                 | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| POST,/api/profiles/otp/verify                                                                               | TRUE            | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          | TRUE          | TRUE         |
| GET,/api/softwares                                                                                          |                 |                 |                 |              |                 | TRUE            |                | TRUE       | TRUE          |               |              |
| GET,/api/softwares/agencies                                                                                 |                 |                 |                 |              |                 | TRUE            |                | TRUE       | TRUE          |               |              |
| GET,/api/softwares/:softwareId                                                                              |                 |                 |                 |              |                 | TRUE            |                | TRUE       | TRUE          |               |              |
| PUT,/api/softwares/:softwareId                                                                              |                 |                 |                 |              |                 | TRUE            |                | TRUE       |               |               |              |
| POST,/api/softwares/upload                                                                                  |                 |                 |                 |              |                 | TRUE            |                | TRUE       |               |               |              |
| POST,/api/softwares/:softwareId/sync                                                                        |                 |                 |                 |              |                 | TRUE            |                | TRUE       |               |               |              |
| POST,/api/softwares/:softwareId/sync-canary                                                                 |                 |                 |                 |              |                 | TRUE            |                | TRUE       |               |               |              |
| POST,/api/softwares/:softwareId/agencies/sync                                                               |                 |                 |                 |              |                 | TRUE            |                | TRUE       |               |               |              |
| POST,/api/softwares/files/verify                                                                            |                 |                 |                 |              |                 | TRUE            |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/feature-layers/:layerId                                                         |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/feature-layers/:layerId/override-info                                          |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/feature-layers/:layerId/rename                                                 |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/feature-layers/upload                                                          |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/feature-layers                                                                  |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| POST,/api/agencies/:agencyId/feature-layers/:layerId/replace                                                |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| DELETE,/api/agencies/:agencyId/feature-layers/:layerId                                                      |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/system-status                                                                                      |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               | TRUE         |
| POST,/api/system-status                                                                                     |                 | TRUE            | TRUE            |              |                 |                 |                | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/devices/preferences/:model                                                      |                 | TRUE            | TRUE            |              | TRUE            |                 | TRUE           | TRUE       | TRUE          |               | TRUE         |
| PUT,/api/agencies/:agencyId/devices/preferences                                                             |                 | TRUE            | TRUE            |              | TRUE            |                 |                | TRUE       | TRUE          |               |              |
| GET,/api/agencies/:agencyId/feature-servers/:layerId                                                        |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/feature-servers/:layerId/query                                                  |                 | TRUE            | TRUE            | TRUE         | TRUE            |                 | TRUE           | TRUE       |               |               |              |
| GET,/api/agencies/:agencyId/applications                                                                    |                 | TRUE            | TRUE            | TRUE         | TRUE            | TRUE            | TRUE           | TRUE       | TRUE          |               | TRUE         |
`

type RoleData struct {
	Endpoint        string
	EDCAUserManager string
	CustomerSupport string
	InstanceManager string
	FieldSupport    string
	ContractSupport string
	FirmwareManager string
	ProductManager  string
	SuperAdmin      string
	DeviceManager   string
	MarketingUser   string
	AgencyViewer    string
}

func Test_serverImpl_getRouter(t *testing.T) {
	type fields struct {
		controller controller.Controller
		middleware middleware.Middleware
	}
	type args struct {
		router   *gin.Engine
		roleData RoleData
	}
	type testCase struct {
		name string
		args args
	}

	tests := []testCase{}
	rolesData := parseMarkdownTable(roleDataMD)
	for _, role := range rolesData {
		tests = append(tests, testCase{
			name: fmt.Sprintf("Test %s", role.Endpoint),
			args: args{
				roleData: role,
			},
		})
	}

	i := &serverImpl{
		controller: controller.NewMockController(t),
		middleware: middleware.NewMockMiddleware(t),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, endpoint := string(strings.Split(tt.args.roleData.Endpoint, ",")[0]), string(strings.Split(tt.args.roleData.Endpoint, ",")[1])

			{
				i.controller.(*controller.MockController).EXPECT().SessionController().Return(
					controller.NewMockSessionController(t),
				)
				i.controller.(*controller.MockController).EXPECT().SoftwareController().Return(
					controller.NewMockSoftwareController(t),
				)
				i.controller.(*controller.MockController).EXPECT().AnnouncementController().Return(
					controller.NewMockAnnouncementController(t),
				)
				i.controller.(*controller.MockController).EXPECT().UtilityController().Return(
					controller.NewMockUtilityController(t),
				)
				i.controller.(*controller.MockController).EXPECT().UserController().Return(
					controller.NewMockUserController(t),
				)
				i.controller.(*controller.MockController).EXPECT().DeviceController().Return(
					controller.NewMockDeviceController(t),
				)
				i.controller.(*controller.MockController).EXPECT().OAuthController().Return(
					controller.NewMockOAuthController(t),
				)
				i.controller.(*controller.MockController).EXPECT().AgencyController().Return(
					controller.NewMockAgencyController(t),
				)
				i.controller.(*controller.MockController).EXPECT().ApplicationController().Return(
					controller.NewMockApplicationController(t),
				)
				i.controller.(*controller.MockController).EXPECT().ReportController().Return(
					controller.NewMockReportController(t),
				)
				i.controller.(*controller.MockController).EXPECT().CadRmsIntegrationController().Return(
					controller.NewMockCadRmsIntegrationController(t),
				)
				i.controller.(*controller.MockController).EXPECT().ProfileController().Return(
					controller.NewMockProfileController(t),
				)
				i.controller.(*controller.MockController).EXPECT().SystemStatusController().Return(
					controller.NewMockSystemStatusController(t),
				)
			}

			valueOfRole := reflect.ValueOf(tt.args.roleData)
			typeOfRole := valueOfRole.Type()
			sessionValidatorHandler := func() gin.HandlerFunc {
				return func(ctx *gin.Context) {
					ctx.Set(constant.Session_Scopes, ctx.Request.Context().Value(constant.Session_Scopes).([]constant.Scope))
					ctx.Next()
				}
			}()

			scopeValidatorHandler := func(scope [][]constant.Scope) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					middleware.CheckScope(ctx, scope)
					err := ctx.Errors.Last()
					if err == nil {
						ctx.Abort()
						return
					}

					var finalErr edcaerror.Error
					switch apiError := err.Err.(type) {
					case edcaerror.Error:
						finalErr = apiError
					default:
						t.Error("Error not found")
					}

					ctx.JSON(finalErr.StatusCode, finalErr)
					ctx.Abort()
				}
			}

			router := i.setupRouter(gin.Default(),
				sessionValidatorHandler,
				scopeValidatorHandler,
			)
			numOfFields := valueOfRole.NumField()
			for index := 0; index < numOfFields; index++ {
				fieldName := strings.Trim(typeOfRole.Field(index).Name, " ")
				fieldValue := strings.Trim(valueOfRole.Field(index).String(), " ")

				if fieldName == "Endpoint" {
					continue
				}

				statusCodeExpected := 200
				if fieldValue != "TRUE" {
					statusCodeExpected = 403
				}

				scopes := getScopesByRole(constant.Role(fieldName))

				w := httptest.NewRecorder()
				req, _ := http.NewRequest(method, endpoint, nil)
				ctx := context.WithValue(req.Context(), constant.Session_Scopes, scopes)
				reqCtx := req.WithContext(ctx)
				router.ServeHTTP(w, reqCtx)
				if w.Code != statusCodeExpected {
					t.Errorf("%s role | %s|%s: Expected status code %d, but got %d", fieldName, method, endpoint, statusCodeExpected, w.Code)
				}
			}
		})
	}
}

func getScopesByRole(role constant.Role) []constant.Scope {
	switch role {
	case constant.SuperAdmin:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.AgencyEDCARead,
			constant.AgencyAnyManage,
			constant.AgencyAnyManageSecurity,
			constant.AgencyAnyManageFeatures,
			constant.AgencyAnyManageIntegrations,
			constant.AgencyCreate,
			constant.UsersAnyRead,
			constant.UsersAnyManage,
			constant.UsersAnyManageSecurity,
			constant.UsersEdcaRead,
			constant.UsersEdcaManage,
			constant.DevicesAnyRead,
			constant.DevicesLogsRead,
			constant.DevicesAnyManageOwner,
			constant.SoftwareRead,
			constant.SoftwareManage,
			constant.AutoTaggingLogsRead,
			constant.AgencyAnyManageAutoTagging,
			constant.AnnouncementsRead,
			constant.AnnouncementsManage,
			constant.StatusRead,
			constant.StatusManage,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
			constant.DeviceAnyExecute,
		}
	case constant.InstanceManager:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.AgencyAnyManage,
			constant.AgencyAnyManageSecurity,
			constant.AgencyAnyManageFeatures,
			constant.AgencyAnyManageIntegrations,
			constant.AgencyCreate,
			constant.UsersAnyRead,
			constant.UsersAnyManage,
			constant.UsersAnyManageSecurity,
			constant.DevicesAnyRead,
			constant.DevicesLogsRead,
			constant.DevicesAnyManageOwner,
			constant.AutoTaggingLogsRead,
			constant.AgencyAnyManageAutoTagging,
			constant.StatusRead,
			constant.StatusManage,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.MarketingUser:
		return []constant.Scope{
			constant.AnnouncementsRead,
			constant.AnnouncementsManage,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.EDCAUserManager:
		return []constant.Scope{
			constant.AgencyEDCARead,
			constant.UsersEdcaRead,
			constant.UsersEdcaManage,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.CustomerSupport:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.AgencyAnyManage,
			constant.AgencyAnyManageSecurity,
			constant.AgencyAnyManageFeatures,
			constant.AgencyAnyManageIntegrations,
			constant.UsersAnyRead,
			constant.UsersAnyManage,
			constant.UsersAnyManageSecurity,
			constant.DevicesAnyRead,
			constant.DevicesLogsRead,
			constant.DevicesAnyManageOwner,
			constant.AutoTaggingLogsRead,
			constant.AgencyAnyManageAutoTagging,
			constant.StatusRead,
			constant.StatusManage,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.FirmwareManager:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.SoftwareRead,
			constant.SoftwareManage,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.ProductManager:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.AgencyAnyManageFeatures,
			constant.UsersAnyRead,
			constant.DevicesAnyRead,
			constant.StatusRead,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.FieldSupport:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.AgencyAnyManage,
			constant.AgencyAnyManageSecurity,
			constant.AgencyAnyManageFeatures,
			constant.AgencyAnyManageIntegrations,
			constant.UsersAnyRead,
			constant.AgencyAnyManageAutoTagging,
			constant.StatusRead,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.ContractSupport:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.AgencyAnyManage,
			constant.AgencyAnyManageFeatures,
			constant.UsersAnyRead,
			constant.UsersAnyManage,
			constant.StatusRead,
			constant.DevicesAnyRead,
			constant.DevicesAnyManageOwner,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.DeviceManager:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.UsersAnyRead,
			constant.DevicesAnyRead,
			constant.DevicesLogsRead,
			constant.DevicesAnyManageOwner,
			constant.SoftwareRead,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	case constant.AgencyViewer:
		return []constant.Scope{
			constant.AgencyAnyRead,
			constant.UsersAnyRead,
			constant.DevicesAnyRead,
			constant.StatusRead,
			constant.UsersSelfRead,
			constant.UsersSelfManage,
		}
	}
	return []constant.Scope{}
}

func parseMarkdownTable(mdText string) []RoleData {
	lines := strings.Split(mdText, "\n")
	var data []RoleData

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
			continue
		}

		cells := strings.Split(strings.Trim(line, "|"), "|")

		for i, cell := range cells {
			cells[i] = strings.TrimSpace(cell)
		}

		if len(data) == 0 {
			data = append(data, RoleData{})
			for _, header := range cells {
				data[0].setField(header, "")
			}
		} else {
			fields := make(map[string]string)

			fieldNames := data[0].getFieldNames()
			for i, header := range fieldNames {
				if i < len(cells) {
					fields[header] = cells[i]
				}
			}

			var roleData RoleData
			mapToStruct(fields, &roleData)

			data = append(data, roleData)
		}
	}

	return data[2:]
}

func (e *RoleData) getFieldNames() []string {
	var fieldNames []string
	t := reflect.TypeOf(*e)
	for i := 0; i < t.NumField(); i++ {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}
	return fieldNames
}

func (e *RoleData) setField(fieldName, value string) {
	v := reflect.ValueOf(e).Elem().FieldByName(fieldName)
	if v.IsValid() {
		v.SetString(value)
	}
}

func mapToStruct(fields map[string]string, result interface{}) {
	resultValue := reflect.ValueOf(result).Elem()

	for fieldName, fieldValue := range fields {
		field := resultValue.FieldByName(fieldName)
		if field.IsValid() {
			field.SetString(fieldValue)
		}
	}
}
