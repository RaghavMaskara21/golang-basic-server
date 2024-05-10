package communities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteCommunityModel struct {
	CommunityId primitive.ObjectID `json:"communityId" validate:"required"`
}

type JoinCommunityModel struct {
	CommunityId primitive.ObjectID `json:"communityId" validate:"required"`
	Password    string             `json:"password"`
}

type LeaveCommunityModel struct {
	CommunityId primitive.ObjectID `json:"communityId" validate:"required"`
}

type RemoveCommunityMembersModel struct {
	CommunityId primitive.ObjectID `json:"communityId" validate:"required"`
	UserId      int32              `json:"userId" validate:"required"`
}

type CommunityAssignModsModel struct {
	CommunityId primitive.ObjectID `json:"communityId" validate:"required"`
	UserId      int32              `json:"userId" validate:"required"`
}

type CommunityFetchInfluencerCommunityModel struct {
	InfluencerId int32 `json:"influencerId" validate:"required"`
}

type GetMyCommunitiesResultModel struct {
	Id               primitive.ObjectID            `json:"id" bson:"_id"`
	CommunityId      primitive.ObjectID            `json:"communityId"`
	IsMod            bool                          `json:"isMod"`
	LastVisited      time.Time                     `json:"lastVisited"`
	IsBanned         bool                          `json:"isBanned"`
	CommunityDetails *CommunityDetailsModel        `json:"communityDetails"`
	MessageDetails   *CommunityMessageDetailsModel `json:"messageDetails"`
	CreatedByName    string                        `json:"createdByName"`
}

type GetCommunityByIdResultModel struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"name"`
	Tags             []string           `json:"tags"`
	Image            string             `json:"image"`
	Bio              string             `json:"bio"`
	Genre            string             `json:"genre"`
	IsPrivate        bool               `json:"isPrivate"`
	TotalCountOfUser int32              `json:"totalCountOfUser"`
	InfluencerId     int32              `json:"influencerId"`
	Priority         int32              `json:"priority"`
	IsJoined         bool               `json:"isJoined"`
	IsMod            bool               `json:"isMod"`
	IsBanned         bool               `json:"isBanned"`
	UserRole         UserRole           `json:"userRole"`
}

type JoinedCommunityHomeResponseModel struct {
	Id                 primitive.ObjectID            `json:"id" bson:"_id"`
	Name               string                        `json:"name"`
	Tags               []string                      `json:"tags"`
	Image              string                        `json:"image"`
	Bio                string                        `json:"bio"`
	Genre              string                        `json:"genre"`
	IsPrivate          bool                          `json:"isPrivate"`
	TotalCountOfUser   int32                         `json:"totalCountOfUser"`
	InfluencerId       int32                         `json:"influencerId"`
	Priority           int32                         `json:"priority"`
	IsMod              bool                          `json:"isMod"`
	IsBanned           bool                          `json:"isBanned"`
	LastVisited        time.Time                     `json:"lastVisited"`
	MessageDetails     *CommunityMessageDetailsModel `json:"messageDetails"`
	UnreadMessageCount int64                         `json:"unreadMessageCount"`
}

type SuggestedCommunityHomeResponseModel struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"name"`
	Tags             []string           `json:"tags"`
	Image            string             `json:"image"`
	Bio              string             `json:"bio"`
	Genre            string             `json:"genre"`
	IsPrivate        bool               `json:"isPrivate"`
	TotalCountOfUser int32              `json:"totalCountOfUser"`
	InfluencerId     int32              `json:"influencerId"`
	Priority         int32              `json:"priority"`
	CreatedByName    string             `json:"createdByName"`
}

type CommunityDetailsModel struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"name"`
	Tags             []string           `json:"tags"`
	Image            string             `json:"image"`
	Bio              string             `json:"bio"`
	Genre            string             `json:"genre"`
	IsPrivate        bool               `json:"isPrivate"`
	TotalCountOfUser int32              `json:"totalCountOfUser"`
	InfluencerId     int32              `json:"influencerId"`
	Priority         int32              `json:"priority"`
}

type CommunityMessageDetailsModel struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	CommunityId primitive.ObjectID `json:"communityId"`
	Message     string             `json:"message"`
	Reactions   interface{}        `json:"reaction"`
	ImageUrl    string             `json:"imageUrl"`
	ByUserId    int32              `json:"byUserId"`
}

type UserCommunityJoinedDetailsModel struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      int32              `json:"userId"`
	CommunityId primitive.ObjectID `json:"communityId"`
	IsMod       bool               `json:"isMod"`
	LastVisited time.Time          `json:"lastVisited"`
	IsBanned    bool               `json:"isBanned"`
}

type UserRole string

const (
	AUDIENCE   UserRole = "AUDIENCE"
	MOD        UserRole = "MOD"
	INFLUENCER UserRole = "INFLUENCER"
)
