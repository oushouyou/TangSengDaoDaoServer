package customersvc

import (
	"fmt"
	"time"

	"github.com/TangSengDaoDao/TangSengDaoDaoServer/modules/group"
	"github.com/TangSengDaoDao/TangSengDaoDaoServer/modules/user"
	"github.com/TangSengDaoDao/TangSengDaoDaoServer/pkg/util"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/common"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/config"
)

type Customer struct {
	Username string
	Phone    string
	Email    string
}

// IService 服务接口
type IService interface {
	//唤起一个客户服务群
	AwakenTheGroup(customer Customer) (string, error)
}

// Service app服务
type Service struct {
	ctx         *config.Context
	groupSvc    group.IService
	userManager user.Manager
	db          *DB
}

// NewService NewService
func NewService(ctx *config.Context) IService {
	return &Service{
		ctx:         ctx,
		groupSvc:    group.NewService(ctx),
		userManager: *user.NewManager(ctx),
		db:          newDB(ctx.DB()),
	}
}

func (s *Service) AwakenTheGroup(customer Customer) (string, error) {

	// 1. 判断客户是否已经有了服务群，如果没有创建，有了直接返回群ID
	cgModel, err := s.db.queryWithCustomerExternalNo(customer.Username)
	if err != nil {
		return "", err
	}
	if cgModel == nil {
		// 1.1 先创建用户。
		// userReq := &user.AddUserReq{
		// 	Name:     customer.Username,
		// 	UID:      "", // 如果无值，则随机生成
		// 	Username: customer.Username,
		// 	Zone:     "biz",
		// 	Phone:    customer.Phone,
		// 	Email:    customer.Email,
		// 	Password: "12345678", // 随机初始密码，真正的登陆，需要提供另外的基于站外JWT的登陆方式。
		// }
		// s.userSvc.AddUser(userReq)
		uid := util.GenerUUID()
		userModel := &user.Model{}
		userModel.UID = uid
		userModel.Name = customer.Username
		userModel.Vercode = fmt.Sprintf("%s@%d", util.GenerUUID(), common.User)
		userModel.QRVercode = fmt.Sprintf("%s@%d", util.GenerUUID(), common.QRCode)
		userModel.Phone = customer.Phone
		userModel.Username = customer.Username
		userModel.Email = customer.Email
		userModel.Zone = "0086"
		userModel.Password = "" //util.MD5(util.MD5("12345678"))
		userModel.ShortNo = util.Ten2Hex(time.Now().UnixNano())
		userModel.IsUploadAvatar = 0
		userModel.NewMsgNotice = 1
		userModel.MsgShowDetail = 1
		userModel.SearchByPhone = 1
		userModel.ShortStatus = 0
		userModel.SearchByShort = 1
		userModel.VoiceOn = 1
		userModel.ShockOn = 1
		userModel.Sex = 1
		userModel.Status = int(common.UserAvailable)
		err = s.userManager.InnerAddUser(userModel)
		if err != nil {
			return "", err
		}
		groupReq := &group.AddGroupReq{
			GroupNo: customer.Username + "_group",
			Name:    "专属服务",
			Creator: s.ctx.GetConfig().Account.SystemUID,
		}
		s.groupSvc.AddGroup(groupReq)
		//客户账号加入群。
		groupMreq := &group.AddMemberReq{
			GroupNo:   groupReq.GroupNo,
			MemberUID: uid,
			Role:      0,
		}
		s.groupSvc.AddMember(groupMreq)
		//系统账号加入群。
		groupMreq2 := &group.AddMemberReq{
			GroupNo:   groupReq.GroupNo,
			MemberUID: s.ctx.GetConfig().Account.SystemUID,
			Role:      1,
		}
		s.groupSvc.AddMember(groupMreq2)
		cgModel = &model{
			CustomerExternalNo: customer.Username,
			UserUid:            uid,
			GroupNo:            groupReq.GroupNo,
		}
		s.db.insert(cgModel)
		// 创建IM频道
		realMemberUids := make([]string, 0) // 真实成员uid集合
		realMemberUids = append(realMemberUids, uid)
		realMemberUids = append(realMemberUids, s.ctx.GetConfig().Account.SystemUID)
		s.ctx.IMCreateOrUpdateChannel(&config.ChannelCreateReq{
			ChannelID:   groupReq.GroupNo,
			ChannelType: common.ChannelTypeGroup.Uint8(),
			Subscribers: realMemberUids,
		})
		return cgModel.GroupNo, nil
	} else {
		return cgModel.GroupNo, nil
	}
}
