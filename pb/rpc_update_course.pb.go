// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: rpc_update_course.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseTitle *string                `protobuf:"bytes,2,opt,name=course_title,json=courseTitle,proto3,oneof" json:"course_title,omitempty"`
	CourseLevel *string                `protobuf:"bytes,3,opt,name=course_level,json=courseLevel,proto3,oneof" json:"course_level,omitempty"`
	StartDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	EndDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	StartTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	EndTime     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	Price       *int64                 `protobuf:"varint,8,opt,name=price,proto3,oneof" json:"price,omitempty"`
	LocationId  *string                `protobuf:"bytes,9,opt,name=location_id,json=locationId,proto3,oneof" json:"location_id,omitempty"`
	MinCapacity *int32                 `protobuf:"varint,10,opt,name=min_capacity,json=minCapacity,proto3,oneof" json:"min_capacity,omitempty"`
	Open        *bool                  `protobuf:"varint,11,opt,name=open,proto3,oneof" json:"open,omitempty"`
	PriceHikeId *string                `protobuf:"bytes,12,opt,name=price_hike_id,json=priceHikeId,proto3,oneof" json:"price_hike_id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	// optional string created_by=15;
	CourseDescription *string `protobuf:"bytes,15,opt,name=course_description,json=courseDescription,proto3,oneof" json:"course_description,omitempty"`
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_rpc_update_course_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCourseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCourseRequest) GetCourseTitle() string {
	if x != nil && x.CourseTitle != nil {
		return *x.CourseTitle
	}
	return ""
}

func (x *UpdateCourseRequest) GetCourseLevel() string {
	if x != nil && x.CourseLevel != nil {
		return *x.CourseLevel
	}
	return ""
}

func (x *UpdateCourseRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *UpdateCourseRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *UpdateCourseRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *UpdateCourseRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *UpdateCourseRequest) GetPrice() int64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *UpdateCourseRequest) GetLocationId() string {
	if x != nil && x.LocationId != nil {
		return *x.LocationId
	}
	return ""
}

func (x *UpdateCourseRequest) GetMinCapacity() int32 {
	if x != nil && x.MinCapacity != nil {
		return *x.MinCapacity
	}
	return 0
}

func (x *UpdateCourseRequest) GetOpen() bool {
	if x != nil && x.Open != nil {
		return *x.Open
	}
	return false
}

func (x *UpdateCourseRequest) GetPriceHikeId() string {
	if x != nil && x.PriceHikeId != nil {
		return *x.PriceHikeId
	}
	return ""
}

func (x *UpdateCourseRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpdateCourseRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpdateCourseRequest) GetCourseDescription() string {
	if x != nil && x.CourseDescription != nil {
		return *x.CourseDescription
	}
	return ""
}

type UpdateCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *UpdateCourseResponse) Reset() {
	*x = UpdateCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseResponse) ProtoMessage() {}

func (x *UpdateCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseResponse.ProtoReflect.Descriptor instead.
func (*UpdateCourseResponse) Descriptor() ([]byte, []int) {
	return file_rpc_update_course_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

var File_rpc_update_course_proto protoreflect.FileDescriptor

var file_rpc_update_course_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x07, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x04, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a,
	0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x08, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x09, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x69, 0x6b, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x48, 0x69,
	0x6b, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x0b, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x0c, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x0d, 0x52, 0x11, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x68, 0x69, 0x6b, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6b, 0x6f, 0x72, 0x6f,
	0x65, 0x6d, 0x65, 0x6b, 0x61, 0x2f, 0x6c, 0x61, 0x2d, 0x63, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x61,
	0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_update_course_proto_rawDescOnce sync.Once
	file_rpc_update_course_proto_rawDescData = file_rpc_update_course_proto_rawDesc
)

func file_rpc_update_course_proto_rawDescGZIP() []byte {
	file_rpc_update_course_proto_rawDescOnce.Do(func() {
		file_rpc_update_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_update_course_proto_rawDescData)
	})
	return file_rpc_update_course_proto_rawDescData
}

var file_rpc_update_course_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_update_course_proto_goTypes = []interface{}{
	(*UpdateCourseRequest)(nil),   // 0: pb.UpdateCourseRequest
	(*UpdateCourseResponse)(nil),  // 1: pb.UpdateCourseResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Course)(nil),                // 3: pb.Course
}
var file_rpc_update_course_proto_depIdxs = []int32{
	2, // 0: pb.UpdateCourseRequest.start_date:type_name -> google.protobuf.Timestamp
	2, // 1: pb.UpdateCourseRequest.end_date:type_name -> google.protobuf.Timestamp
	2, // 2: pb.UpdateCourseRequest.start_time:type_name -> google.protobuf.Timestamp
	2, // 3: pb.UpdateCourseRequest.end_time:type_name -> google.protobuf.Timestamp
	2, // 4: pb.UpdateCourseRequest.created_at:type_name -> google.protobuf.Timestamp
	2, // 5: pb.UpdateCourseRequest.updated_at:type_name -> google.protobuf.Timestamp
	3, // 6: pb.UpdateCourseResponse.course:type_name -> pb.Course
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_rpc_update_course_proto_init() }
func file_rpc_update_course_proto_init() {
	if File_rpc_update_course_proto != nil {
		return
	}
	file_course_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_update_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_update_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_rpc_update_course_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_update_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_update_course_proto_goTypes,
		DependencyIndexes: file_rpc_update_course_proto_depIdxs,
		MessageInfos:      file_rpc_update_course_proto_msgTypes,
	}.Build()
	File_rpc_update_course_proto = out.File
	file_rpc_update_course_proto_rawDesc = nil
	file_rpc_update_course_proto_goTypes = nil
	file_rpc_update_course_proto_depIdxs = nil
}