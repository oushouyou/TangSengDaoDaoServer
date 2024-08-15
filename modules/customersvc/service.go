package customersvc

import (
	"github.com/TangSengDaoDao/TangSengDaoDaoServer/modules/group"
	"github.com/TangSengDaoDao/TangSengDaoDaoServer/modules/user"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/config"
)

// IService 服务接口
type IService interface {
	//唤起一个客户服务群
	AwakenTheGroup(externalNo string) (string, error)
}

// Service app服务
type Service struct {
	ctx      *config.Context
	userSvc  user.IService
	groupSvc group.IService
	db       *DB
}

// NewService NewService
func NewService(ctx *config.Context) IService {
	return &Service{
		ctx:      ctx,
		userSvc:  user.NewService(ctx),
		groupSvc: group.NewService(ctx),
		db:       newDB(ctx.DB()),
	}
}

func (s *Service) AwakenTheGroup(externalNo string) (string, error) {
	// 1. 判断客户是否已经有了服务群，如果没有创建，有了直接返回群ID
	cgModel, err := s.db.queryWithCustomerExternalNo(externalNo)

	if err != nil {
		return "", err
	}
	if cgModel == nil {
		// 1.1 先创建用户。
		userReq := &user.AddUserReq{
			Name:     "string",
			UID:      "string", // 如果无值，则随机生成
			Username: "string",
			Zone:     "string",
			Phone:    "string",
			Email:    "string",
			Password: "string", // 随机初始密码，真正的登陆，需要提供另外的基于站外JWT的登陆方式。
		}
		s.userSvc.AddUser(userReq)
		groupReq := &group.AddGroupReq{
			GroupNo: "",
			Name:    "string",
		}
		s.groupSvc.AddGroup(groupReq)
		groupMreq := &group.AddMemberReq{
			GroupNo:   "string",
			MemberUID: "",
		}
		s.groupSvc.AddMember(groupMreq)
		cgModel = &model{
			CustomerExternalNo: externalNo,
			// UserId             string
			// GroupId            string
		}
		s.db.insert(cgModel)
		return "", nil
	} else {
		return cgModel.GroupId, nil
	}
}
