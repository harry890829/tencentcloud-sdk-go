// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180717

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AiAnalysisResult struct {

	// 任务的类型，可以取的值有：
	// <li>Classification：智能分类</li>
	// <li>Cover：智能封面</li>
	// <li>Description：智能描述</li>
	// <li>Highlight：智能精彩片断</li>
	// <li>Tag：智能标签</li>
	// <li>FrameTag：智能按帧标签</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频内容分析智能分类任务的查询结果，当任务类型为 Classification 时有效。
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitempty" name:"ClassificationTask"`

	// 视频内容分析智能封面任务的查询结果，当任务类型为 Cover 时有效。
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitempty" name:"CoverTask"`

	// 视频内容分析智能标签任务的查询结果，当任务类型为 Tag 时有效。
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitempty" name:"TagTask"`
}

type AiAnalysisTaskClassificationInput struct {

	// 视频智能分类模板 ID ，固定为 10 。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {

	// 视频智能分类列表。
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitempty" name:"ClassificationSet" list`
}

type AiAnalysisTaskClassificationResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能分类任务输入。
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitempty" name:"Input"`

	// 智能分类任务输出。
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskCoverInput struct {

	// 视频智能封面模板 ID ，固定为 10 。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {

	// 智能封面列表。
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitempty" name:"CoverSet" list`
}

type AiAnalysisTaskCoverResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能封面任务输入。
	Input *AiAnalysisTaskCoverInput `json:"Input,omitempty" name:"Input"`

	// 智能分类任务输出。
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskInput struct {

	// 视频内容分析模板 ID ，固定为 10，同时进行智能分类、智能标签分析、智能封面分析。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagInput struct {

	// 视频智能分类模板 ID ，固定为 10 。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagOutput struct {

	// 视频智能标签列表。
	TagSet *MediaAiAnalysisTagItem `json:"TagSet,omitempty" name:"TagSet"`
}

type AiAnalysisTaskTagResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能标签任务输入。
	Input *AiAnalysisTaskTagInput `json:"Input,omitempty" name:"Input"`

	// 智能标签任务输出。
	Output []*AiAnalysisTaskTagOutput `json:"Output,omitempty" name:"Output" list`
}

type AiContentReviewResult struct {

	// 任务的类型，可以取的值有：
	// <li>Porn：图片鉴黄</li>
	// <li>Terrorism：图片鉴恐</li>
	// <li>Political：图片鉴政</li>
	// <li>Porn.Asr：Asr 文字鉴黄</li>
	// <li>Porn.Ocr：Ocr 文字鉴黄</li>
	// <li>Political.Asr：Asr 文字鉴政</li>
	// <li>Political.Ocr：Ocr 文字鉴政</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频内容审核智能画面鉴黄任务的查询结果，当任务类型为 Porn 时有效。
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitempty" name:"PornTask"`

	// 视频内容审核智能画面鉴恐任务的查询结果，当任务类型为 Terrorism 时有效。
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitempty" name:"TerrorismTask"`

	// 视频内容审核智能画面鉴恐任务的查询结果，当任务类型为 Political 时有效。
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitempty" name:"PoliticalTask"`

	// 视频内容审核 Asr 文字鉴黄任务的查询结果，当任务类型为 Porn.Asr 时有效。
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitempty" name:"PornAsrTask"`

	// 视频内容审核 Ocr 文字鉴黄任务的查询结果，当任务类型为 Porn.Ocr 时有效。
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitempty" name:"PornOcrTask"`

	// 视频内容审核 Asr 文字鉴政任务的查询结果，当任务类型为 Political.Asr 时有效。
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitempty" name:"PoliticalAsrTask"`

	// 视频内容审核 Ocr 文字鉴政任务的查询结果，当任务类型为 Political.Ocr 时有效。
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitempty" name:"PoliticalOcrTask"`
}

type AiContentReviewTaskInput struct {

	// 视频内容审核模板 ID，可以填 10 和 20。<li>10：对视频画面进行鉴黄、鉴暴、鉴政审核；</li><li>20：对视频画面进行鉴黄、鉴暴、鉴政审核，并对 Asr 和 Ocr 文字进行鉴黄、鉴政审核。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskInput struct {

	// 鉴政模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {

	// Asr 文字涉政、敏感评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Asr 文字涉政、敏感结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Asr 文字有涉政、敏感嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPoliticalOcrTaskInput struct {

	// 鉴政模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {

	// Ocr 文字涉政、敏感评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Ocr 文字涉政、敏感结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Ocr 文字有涉政、敏感嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPoliticalTaskInput struct {

	// 鉴政模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {

	// 视频涉政评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 涉政结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频鉴政结果标签，取值范围：
	// <li>politician：政治人物。</li>
	// <li>violation_photo：违规图标。</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有涉政嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornAsrTaskInput struct {

	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {

	// Asr 文字涉黄评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Asr 文字涉黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Asr 文字有涉黄嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornOcrTaskInput struct {

	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {

	// Ocr 文字涉黄评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Ocr 文字涉黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Ocr 文字有涉黄嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornTaskInput struct {

	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {

	// 视频鉴黄评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 鉴黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频鉴黄结果标签，取值范围：
	// <li>porn：色情。</li>
	// <li>sexy：性感。</li>
	// <li>vulgar：低俗。</li>
	// <li>intimacy：亲密行为。</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有涉黄嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewTaskPoliticalAsrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Asr 文字鉴政任务输入。
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Asr 文字鉴政任务输出。
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Ocr 文字鉴政任务输入。
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Ocr 文字鉴政任务输出。
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核鉴政任务输入。
	Input *AiReviewPoliticalTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核鉴政任务输出。
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Asr 文字鉴黄任务输入。
	Input *AiReviewPornAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Asr 文字鉴黄任务输出。
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Ocr 文字鉴黄任务输入。
	Input *AiReviewPornOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Ocr 文字鉴黄任务输出。
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核鉴黄任务输入。
	Input *AiReviewPornTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核鉴黄任务输出。
	Output *AiReviewPornTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核鉴恐任务输入。
	Input *AiReviewTerrorismTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核鉴恐任务输出。
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTerrorismTaskInput struct {

	// 鉴恐模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {

	// 视频暴恐评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 暴恐结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频暴恐结果标签，取值范围：
	// <li>guns：武器枪支。</li>
	// <li>crowd：人群聚集。</li>
	// <li>police：警察部队。</li>
	// <li>bloody：血腥画面。</li>
	// <li>banners：暴恐旗帜。</li>
	// <li>militant：武装分子。</li>
	// <li>explosion：爆炸火灾。</li>
	// <li>terrorists：暴恐人物。</li>
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有暴恐嫌疑的视频片段列表。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AnimatedGraphicTaskInput struct {

	// 视频转动图模板 ID
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 动图在视频中的开始时间，单位为秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间，单位为秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type ApplyUploadRequest struct {
	*tchttp.BaseRequest

	// 媒体类型，可选值请参考[上传能力综述](https://cloud.tencent.com/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	MediaType *string `json:"MediaType,omitempty" name:"MediaType"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// 封面类型，可选值请参考[上传能力综述](https://cloud.tencent.com/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	CoverType *string `json:"CoverType,omitempty" name:"CoverType"`

	// 媒体后续任务操作，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户。
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过[创建分类](https://cloud.tencent.com/document/product/266/7812)接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，上传回调接口将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ApplyUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储桶，用于上传接口 URL 的 bucket_name。
		StorageBucket *string `json:"StorageBucket,omitempty" name:"StorageBucket"`

		// 存储园区，用于上传接口 Host 的 Region。
		StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

		// 点播会话，用于确认上传接口的参数 VodSessionKey。
		VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

		// 媒体存储路径，用于上传接口存储媒体的对象键（Key）。
		MediaStoragePath *string `json:"MediaStoragePath,omitempty" name:"MediaStoragePath"`

		// 封面存储路径，用于上传接口存储封面的对象键（Key）。
		CoverStoragePath *string `json:"CoverStoragePath,omitempty" name:"CoverStoragePath"`

		// 临时凭证，用于上传接口的权限验证。
		TempCertificate *TempCertificate `json:"TempCertificate,omitempty" name:"TempCertificate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClipFileInfo2017 struct {

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误描述。
	Message *float64 `json:"Message,omitempty" name:"Message"`

	// 输出目标文件的文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 输出目标文件的文件地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 输出目标文件的文件类型。
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ClipTask2017 struct {

	// 视频剪辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 视频剪辑任务源文件 ID。
	SrcFileId *string `json:"SrcFileId,omitempty" name:"SrcFileId"`

	// 视频剪辑输出的文件信息。
	FileInfo *ClipFileInfo2017 `json:"FileInfo,omitempty" name:"FileInfo"`
}

type CommitUploadRequest struct {
	*tchttp.BaseRequest

	// 点播会话，取申请上传接口的返回值 VodSessionKey。
	VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CommitUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体文件的唯一标识。
		FileId *string `json:"FileId,omitempty" name:"FileId"`

		// 媒体播放地址。
		MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

		// 媒体封面地址。
		CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConcatFileInfo2017 struct {

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频拼接源文件的 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 视频拼接源文件的地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 视频拼接源文件的格式。
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ConcatTask2017 struct {

	// 视频拼接任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 视频拼接源文件信息。
	FileInfoSet []*ConcatFileInfo2017 `json:"FileInfoSet,omitempty" name:"FileInfoSet" list`
}

type ConfirmEventsRequest struct {
	*tchttp.BaseRequest

	// 事件句柄，数组长度限制：16。
	EventHandles []*string `json:"EventHandles,omitempty" name:"EventHandles" list`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ConfirmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConfirmEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CoverBySnapshotTaskInput struct {

	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截图方式。包含：
	// <li>Time：依照时间点截图</li>
	// <li>Percent：依照百分比截图</li>
	PositionType *string `json:"PositionType,omitempty" name:"PositionType"`

	// 截图位置：
	// <li>对于依照时间点截图，该值表示指定视频第几秒的截图作为封面</li>
	// <li>对于依照百分比截图，该值表示使用视频百分之多少的截图作为封面</li>
	PositionValue *float64 `json:"PositionValue,omitempty" name:"PositionValue"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type CoverBySnapshotTaskOutput struct {

	// 封面 URL。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`
}

type CreateClassRequest struct {
	*tchttp.BaseRequest

	// 父类 ID，一级分类填写 -1。
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 分类名称，长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类 ID
		ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTask2017 struct {

	// 截图雪碧图任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 截取雪碧图文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 雪碧图规格，参见[雪碧图截图模板](https://cloud.tencent.com/document/product/266/11702#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 雪碧图小图总数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 截取雪碧图输出的地址。
	ImageSpriteUrlSet []*string `json:"ImageSpriteUrlSet,omitempty" name:"ImageSpriteUrlSet" list`

	// 雪碧图子图位置与时间关系 WebVtt 文件地址。
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type DeleteClassRequest struct {
	*tchttp.BaseRequest

	// 分类 ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest

	// 媒体文件的唯一标识。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 指定本次需要删除的部分。默认值为 "[]", 表示删除媒体及其对应的全部视频处理文件。
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitempty" name:"DeleteParts" list`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllClassRequest struct {
	*tchttp.BaseRequest

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAllClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类信息集合
		ClassInfoSet []*MediaClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID 列表，N 从 0 开始取值，最大 19。
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds" list`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	Filters []*string `json:"Filters,omitempty" name:"Filters" list`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeMediaInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体文件信息列表。
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet" list`

		// 不存在的文件 ID 列表。
		NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitempty" name:"NotExistFileIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 视频处理任务的任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务类型，取值：
	// <li>Procedure：视频处理任务；</li>
	// <li>EditMedia：视频编辑任务；</li>
	// <li>WechatPublish：微信发布任务。</li>
	// 兼容 2017 版的任务类型：
	// <li>Transcode：视频转码任务；</li>
	// <li>SnapshotByTimeOffset：视频截图任务；</li>
	// <li>Concat：视频拼接任务；</li>
	// <li>Clip：视频剪辑任务；</li>
	// <li>ImageSprites：截取雪碧图任务。</li>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 任务状态，取值：
	// <li>WAITING：等待中；</li>
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务的创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

		// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

		// 视频处理任务信息，仅当 TaskType 为 Procedure，该字段有值。
		ProcedureTask *ProcedureTask `json:"ProcedureTask,omitempty" name:"ProcedureTask"`

		// 视频编辑任务信息，仅当 TaskType 为 EditMedia，该字段有值。
		EditMediaTask *EditMediaTask `json:"EditMediaTask,omitempty" name:"EditMediaTask"`

		// 微信发布任务信息，仅当 TaskType 为 WechatPublish，该字段有值。
		WechatPublishTask *WechatPublishTask `json:"WechatPublishTask,omitempty" name:"WechatPublishTask"`

		// 视频转码任务信息，仅当 TaskType 为 Transcode，该字段有值。
		TranscodeTask *TranscodeTask2017 `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

		// 视频指定时间点截图任务信息，仅当 TaskType 为 SnapshotByTimeOffset，该字段有值。
		SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

		// 视频拼接任务信息，仅当 TaskType 为 Concat，该字段有值。
		ConcatTask *ConcatTask2017 `json:"ConcatTask,omitempty" name:"ConcatTask"`

		// 视频剪辑任务信息，仅当 TaskType 为 Clip，该字段有值。
		ClipTask *ClipTask2017 `json:"ClipTask,omitempty" name:"ClipTask"`

		// 截取雪碧图任务信息，仅当 TaskType 为 ImageSprite，该字段有值。
		CreateImageSpriteTask *CreateImageSpriteTask2017 `json:"CreateImageSpriteTask,omitempty" name:"CreateImageSpriteTask"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 过滤条件：文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务概要列表。
		TaskSet []*TaskSimpleInfo `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 ID。当该字段为空，说明已无更多数据。
		ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaFileInfo struct {

	// 视频的 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 视频剪辑的起始偏移时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 视频剪辑的起始结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type EditMediaStreamInfo struct {

	// 录制的流 ID
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 流剪辑的起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type EditMediaTask struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频编辑任务的输入。
	Input *EditMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// 视频编辑任务的输出。
	Output *EditMediaTaskOutput `json:"Output,omitempty" name:"Output"`

	// 若发起视频编辑任务时指定了视频处理流程，则该字段为流程任务 ID。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`
}

type EditMediaTaskInput struct {

	// 输入视频的来源类型，可以取的值为 File，Stream 两种。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 输入的视频文件信息，当 InputType 为 File 时，该字段有值。
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitempty" name:"FileInfoSet" list`

	// 输入的流信息，当 InputType 为 Stream 时，该字段有值。
	StreamInfoSet []*EditMediaStreamInfo `json:"StreamInfoSet,omitempty" name:"StreamInfoSet" list`
}

type EditMediaTaskOutput struct {

	// 文件类型，例如 mp4、flv 等。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

type EventContent struct {

	// 事件句柄，调用方必须调用 ConfirmEvents 来确认消息已经收到，确认有效时间 30 秒。失效后，事件可重新被获取。
	EventHandle *string `json:"EventHandle,omitempty" name:"EventHandle"`

	// <b>支持事件类型：</b>
	// <li>NewFileUpload：视频上传完成；</li>
	// <li>ProcedureStateChanged：任务流状态变更；</li>
	// <li>FileDeleted：视频删除完成；</li>
	// <li>PullComplete：视频转拉完成；</li>
	// <li>EditMediaComplete：视频编辑完成；</li>
	// <li>WechatPublishComplete：微信发布完成。</li>
	// <b>兼容 2017 版的事件类型：</b>
	// <li>TranscodeComplete：视频转码完成；</li>
	// <li>ConcatComplete：视频拼接完成；</li>
	// <li>ClipComplete：视频剪辑完成；</li>
	// <li>CreateImageSpriteComplete：视频截取雪碧图完成；</li>
	// <li>CreateSnapshotByTimeOffsetComplete：视频按时间点截图完成。</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 视频上传完成事件，当事件类型为 NewFileUpload 时有效。
	FileUploadEvent *FileUploadTask `json:"FileUploadEvent,omitempty" name:"FileUploadEvent"`

	// 任务流状态变更事件，当事件类型为 ProcedureStateChanged 时有效。
	ProcedureStateChangeEvent *ProcedureTask `json:"ProcedureStateChangeEvent,omitempty" name:"ProcedureStateChangeEvent"`

	// 文件删除事件，当事件类型为 FileDeleted 时有效。
	FileDeleteEvent *FileDeleteTask `json:"FileDeleteEvent,omitempty" name:"FileDeleteEvent"`

	// 视频转拉完成事件，当事件类型为 PullComplete 时有效。
	PullCompleteEvent *PullFileTask `json:"PullCompleteEvent,omitempty" name:"PullCompleteEvent"`

	// 视频编辑完成事件，当事件类型为 EditMediaComplete 时有效。
	EditMediaCompleteEvent *EditMediaTask `json:"EditMediaCompleteEvent,omitempty" name:"EditMediaCompleteEvent"`

	// 微信发布完成事件，当事件类型为 WechatPublishComplete 时有效。
	WechatPublishComplete *WechatPublishTask `json:"WechatPublishComplete,omitempty" name:"WechatPublishComplete"`

	// 视频转码完成事件，当事件类型为 TranscodeComplete 时有效。
	TranscodeCompleteEvent *TranscodeTask2017 `json:"TranscodeCompleteEvent,omitempty" name:"TranscodeCompleteEvent"`

	// 视频拼接完成事件，当事件类型为 ConcatComplete 时有效。
	ConcatCompleteEvent *ConcatTask2017 `json:"ConcatCompleteEvent,omitempty" name:"ConcatCompleteEvent"`

	// 视频剪辑完成事件，当事件类型为 ClipComplete 时有效。
	ClipCompleteEvent *ClipTask2017 `json:"ClipCompleteEvent,omitempty" name:"ClipCompleteEvent"`

	// 视频截取雪碧图完成事件，当事件类型为 CreateImageSpriteComplete 时有效。
	CreateImageSpriteCompleteEvent *CreateImageSpriteTask2017 `json:"CreateImageSpriteCompleteEvent,omitempty" name:"CreateImageSpriteCompleteEvent"`

	// 视频按时间点截图完成事件，当事件类型为 CreateSnapshotByTimeOffsetComplete 时有效。
	SnapshotByTimeOffsetCompleteEvent *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetCompleteEvent,omitempty" name:"SnapshotByTimeOffsetCompleteEvent"`
}

type FileDeleteTask struct {

	// 删除文件 ID 列表。
	FileIdSet []*string `json:"FileIdSet,omitempty" name:"FileIdSet" list`
}

type FileUploadTask struct {

	// 文件唯一 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 上传完成后生成的媒体文件基础信息。
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitempty" name:"MediaBasicInfo"`

	// 若视频上传时指定了视频处理流程，则该字段为流程任务 ID。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`
}

type ImageSpriteTaskInput struct {

	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type LiveRealTimeClipRequest struct {
	*tchttp.BaseRequest

	// 推流[直播码](https://cloud.tencent.com/document/product/267/5959)。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 流剪辑的开始时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否固化。0 不固化，1 固化。默认不固化。
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`

	// 剪辑固化后的视频存储过期时间。格式参照 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。仅 IsPersistence 为 1 时有效，默认剪辑固化的视频永不过期。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 剪辑固化后的视频点播任务流处理，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。仅 IsPersistence 为 1 时有效。
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`
}

func (r *LiveRealTimeClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveRealTimeClipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveRealTimeClipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 剪辑后的视频播放 URL。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 剪辑固化后的视频的媒体文件的唯一标识。
		FileId *string `json:"FileId,omitempty" name:"FileId"`

		// 剪辑固化后的视频任务流 ID。
		VodTaskId *string `json:"VodTaskId,omitempty" name:"VodTaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveRealTimeClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveRealTimeClipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MediaAiAnalysisClassificationItem struct {

	// 智能分类的类别名称。
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// 智能分类的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {

	// 智能封面地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 智能封面的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisTagItem struct {

	// 标签名称。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAnimatedGraphicsInfo struct {

	// 视频转动图结果信息
	AnimatedGraphicsSet []*MediaAnimatedGraphicsItem `json:"AnimatedGraphicsSet,omitempty" name:"AnimatedGraphicsSet" list`
}

type MediaAnimatedGraphicsItem struct {

	// 转动图的文件地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/266/11701#.E9.A2.84.E7.BD.AE.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 动图格式，如 gif。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 动图的高度，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 动图的宽度，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 动图码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 动图大小，单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 动图的md5值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {

	// 音频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	SamplingRate *int64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type MediaBasicInfo struct {

	// 媒体文件名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体文件描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 媒体文件的创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 媒体文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 媒体文件的过期时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。“9999-12-31T23:59:59Z”表示永不过期。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 媒体文件的分类 ID。
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 媒体文件的分类名称。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 媒体文件的分类路径，分类间以“-”分隔，如“新的一级分类 - 新的二级分类”。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 媒体文件的封装格式，例如 mp4、flv 等。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 原始媒体文件的 URL 地址。
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// 该媒体文件的来源信息。
	SourceInfo *MediaSourceData `json:"SourceInfo,omitempty" name:"SourceInfo"`

	// 媒体文件存储地区，如 ap-guangzhou，参见[地域列表](https://cloud.tencent.com/document/api/213/15692#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)。
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 媒体文件的标签信息。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 直播录制文件的唯一标识
	Vid *string `json:"Vid,omitempty" name:"Vid"`
}

type MediaClassInfo struct {

	// 分类 ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 父类 ID，一级分类的父类 ID 为 -1。
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 分类名称
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 分类级别，一级分类为 0，最大值为 3，即最多允许 4 级分类层次。
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 当前分类的第一级子类 ID 集合
	SubClassIdSet []*int64 `json:"SubClassIdSet,omitempty" name:"SubClassIdSet" list`
}

type MediaContentReviewAsrTextSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet" list`
}

type MediaContentReviewOcrTextSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet" list`

	// 嫌疑文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type MediaContentReviewPoliticalSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉政分数。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段鉴政结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 涉政人物、违规图标名字。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 嫌疑片段鉴政结果标签。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 涉政人物、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`
}

type MediaContentReviewSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉黄分数。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段鉴黄结果标签。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`
}

type MediaDeleteItem struct {

	// 所指定的删除部分。如果未填写该字段则参数无效。可选值有：
	// <li>TranscodeFiles（删除转码文件）。</li>
	// <li>WechatPublishFiles（删除微信发布文件）。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 删除由Type参数指定的种类下的视频模板号，模板定义参见[转码模板](https://cloud.tencent.com/document/product/266/11701#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	// 默认值为0，表示删除参数Type指定种类下所有的视频。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type MediaImageSpriteInfo struct {

	// 特定规格的雪碧图信息集合，每个元素代表一套相同规格的雪碧图。
	ImageSpriteSet []*MediaImageSpriteItem `json:"ImageSpriteSet,omitempty" name:"ImageSpriteSet" list`
}

type MediaImageSpriteItem struct {

	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/11702#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 雪碧图小图的高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 每一张雪碧图大图的地址。
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// 雪碧图子图位置与时间关系的 WebVtt 文件地址。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type MediaInfo struct {

	// 基础信息。包括视频名称、大小、时长、封面图片等。
	BasicInfo *MediaBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 元信息。包括视频流信息、音频流信息等。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 转码结果信息。包括该视频转码生成的各种码率的视频的地址、规格、码率、分辨率等。
	TranscodeInfo *MediaTranscodeInfo `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// 转动图结果信息。对视频转动图（如 gif）后，动图相关信息。
	AnimatedGraphicsInfo *MediaAnimatedGraphicsInfo `json:"AnimatedGraphicsInfo,omitempty" name:"AnimatedGraphicsInfo"`

	// 采样截图信息。对视频采样截图后，相关截图信息。
	SampleSnapshotInfo *MediaSampleSnapshotInfo `json:"SampleSnapshotInfo,omitempty" name:"SampleSnapshotInfo"`

	// 雪碧图信息。对视频截取雪碧图之后，雪碧的相关信息。
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo,omitempty" name:"ImageSpriteInfo"`

	// 指定时间点截图信息。对视频依照指定时间点截图后，各个截图的信息。
	SnapshotByTimeOffsetInfo *MediaSnapshotByTimeOffsetInfo `json:"SnapshotByTimeOffsetInfo,omitempty" name:"SnapshotByTimeOffsetInfo"`

	// 视频打点信息。对视频设置的各个打点信息。
	KeyFrameDescInfo *MediaKeyFrameDescInfo `json:"KeyFrameDescInfo,omitempty" name:"KeyFrameDescInfo"`

	// 媒体文件唯一标识 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type MediaInputInfo struct {

	// 视频 URL。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 视频名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频自定义 ID。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type MediaKeyFrameDescInfo struct {

	// 视频打点信息数组。
	KeyFrameDescSet []*MediaKeyFrameDescItem `json:"KeyFrameDescSet,omitempty" name:"KeyFrameDescSet" list`
}

type MediaKeyFrameDescItem struct {

	// 打点的视频偏移时间，单位：秒。
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 打点的内容字符串，限制 1-128 个字符。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type MediaMetaData struct {

	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// 视频时长，单位：秒。
	VideoDuration *float64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// 音频时长，单位：秒。
	AudioDuration *float64 `json:"AudioDuration,omitempty" name:"AudioDuration"`
}

type MediaOutputInfo struct {

	// 输出文件 Bucket 所属地域，如 ap-guangzhou  。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 输出文件 Bucket 。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 输出文件目录，目录名必须以 "/" 结尾。
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type MediaProcessTaskAnimatedGraphicResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转动图任务的输入。
	Input *AnimatedGraphicTaskInput `json:"Input,omitempty" name:"Input"`

	// 转动图任务的输出。
	Output *MediaAnimatedGraphicsItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskCoverBySnapshotResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频截图做封面任务的输入。
	Input *CoverBySnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频截图做封面任务的输出。
	Output *CoverBySnapshotTaskOutput `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频截雪碧图任务的输入。
	Input *ImageSpriteTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频截雪碧图任务的输出。
	Output *MediaImageSpriteItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskInput struct {

	// 视频转码任务列表。
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet" list`

	// 视频转动图任务列表。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet" list`

	// 对视频按时间点截图任务列表。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet" list`

	// 对视频采样截图任务列表。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet" list`

	// 对视频截雪碧图任务列表。
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet" list`

	// 对视频截图做封面任务列表。
	CoverBySnapshotTaskSet []*CoverBySnapshotTaskInput `json:"CoverBySnapshotTaskSet,omitempty" name:"CoverBySnapshotTaskSet" list`
}

type MediaProcessTaskResult struct {

	// 任务的类型，可以取的值有：
	// <li>Transcode：转码</li>
	// <li>AnimatedGraphics：转动图</li>
	// <li>SnapshotByTimeOffset：时间点截图</li>
	// <li>SampleSnapshot：采样截图</li>
	// <li>ImageSprites：雪碧图</li>
	// <li>CoverBySnapshot：截图做封面</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

	// 视频转动图任务的查询结果，当任务类型为 AnimatedGraphics 时有效。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitempty" name:"AnimatedGraphicTask"`

	// 对视频按时间点截图任务的查询结果，当任务类型为 SnapshotByTimeOffset 时有效。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

	// 对视频采样截图任务的查询结果，当任务类型为 SampleSnapshot 时有效。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitempty" name:"SampleSnapshotTask"`

	// 对视频截雪碧图任务的查询结果，当任务类型为 ImageSprite 时有效。
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitempty" name:"ImageSpriteTask"`

	// 对视频截图做封面任务的查询结果，当任务类型为 CoverBySnapshot 时有效。
	CoverBySnapshotTask *MediaProcessTaskCoverBySnapshotResult `json:"CoverBySnapshotTask,omitempty" name:"CoverBySnapshotTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频做采样截图任务输入。
	Input *SampleSnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频做采样截图任务输出。
	Output *MediaSampleSnapshotItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频按指定时间点截图任务输入。
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频按指定时间点截图任务输出。
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskTranscodeResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInput `json:"Input,omitempty" name:"Input"`

	// 转码任务的输出。
	Output *MediaTranscodeItem `json:"Output,omitempty" name:"Output"`
}

type MediaSampleSnapshotInfo struct {

	// 特定规格的采样截图信息集合，每个元素代表一套相同规格的采样截图。
	SampleSnapshotSet []*MediaSampleSnapshotItem `json:"SampleSnapshotSet,omitempty" name:"SampleSnapshotSet" list`
}

type MediaSampleSnapshotItem struct {

	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/11702#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 生成的截图 url 列表。
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSnapshotByTimeOffsetInfo struct {

	// 特定规格的指定时间点截图信息集合。目前每种规格只能有一套截图。
	SnapshotByTimeOffsetSet []*MediaSnapshotByTimeOffsetItem `json:"SnapshotByTimeOffsetSet,omitempty" name:"SnapshotByTimeOffsetSet" list`
}

type MediaSnapshotByTimeOffsetItem struct {

	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/11702#.E6.8C.87.E5.AE.9A.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitempty" name:"PicInfoSet" list`
}

type MediaSnapshotByTimePicInfoItem struct {

	// 该张截图对应视频文件中的时间偏移，单位：秒。
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 该张截图的 URL 地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSourceData struct {

	// 媒体文件的来源类别：
	// <li>Record：来自录制。如直播录制、直播时移录制等。</li>
	// <li>Upload：来自上传。如拉取上传、服务端上传、客户端 UGC 上传等。</li>
	// <li>VideoProcessing：来自视频处理。如视频拼接、视频剪辑等。</li>
	// <li>Unknown：未知来源。</li>
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 用户创建文件时透传的字段
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type MediaTranscodeInfo struct {

	// 各规格的转码信息集合，每个元素代表一个规格的转码结果。
	TranscodeSet []*MediaTranscodeItem `json:"TranscodeSet,omitempty" name:"TranscodeSet" list`
}

type MediaTranscodeItem struct {

	// 转码后的视频文件地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/11701#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 媒体文件总大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频的 md5 值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`
}

type MediaVideoStreamItem struct {

	// 视频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流的高度，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流的宽度，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频流的编码格式，例如 h264。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}

type ModifyClassRequest struct {
	*tchttp.BaseRequest

	// 分类 ID
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// 分类名称。长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaInfoRequest struct {
	*tchttp.BaseRequest

	// 媒体文件唯一标识。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件名称，最长 64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体文件描述，最长 128 个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 媒体文件分类 ID。
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 媒体文件过期时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 视频封面图片文件（如 jpeg, png 等）进行 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 新增的一组视频打点信息，如果某个偏移时间已存在打点，则会进行覆盖操作，单个媒体文件最多 100 个打点信息。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs,omitempty" name:"AddKeyFrameDescs" list`

	// 要删除的一组视频打点信息的时间偏移，单位：秒。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs,omitempty" name:"DeleteKeyFrameDescs" list`

	// 取值 1 表示清空视频打点信息，其他值无意义。
	// 同一个请求里，ClearKeyFrameDescs 与 AddKeyFrameDescs 不能同时出现。
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs,omitempty" name:"ClearKeyFrameDescs"`

	// 新增的一组标签，单个媒体文件最多 16 个标签，单个标签最多 16 个字符。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	AddTags []*string `json:"AddTags,omitempty" name:"AddTags" list`

	// 要删除的一组标签。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	DeleteTags []*string `json:"DeleteTags,omitempty" name:"DeleteTags" list`

	// 取值 1 表示清空媒体文件所有标签，其他值无意义。
	// 同一个请求里，ClearTags 与 AddTags 不能同时出现。
	ClearTags *int64 `json:"ClearTags,omitempty" name:"ClearTags"`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyMediaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新的视频封面 URL。
	// * 注意：仅当请求携带 CoverData 时此返回值有效。 *
		CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcedureTask struct {

	// 视频处理任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 媒体文件 ID
	// <li>若流程由 [ProcessMedia](http://139.199.214.26/document/product/266/30248) 发起，该字段表示 [MediaInfo](http://139.199.214.26/document/product/266/16694?!preview&preview_docmenu=1&lang=cn&!document=1#MediaInfo) 的 FileId；</li>
	// <li>若流程由 [ProcessMediaByUrl](http://139.199.214.26/document/product/266/31058?!document=1&!preview) 发起，该字段表示 [MediaInputInfo](http://139.199.214.26/document/product/266/16694?!preview&preview_docmenu=1&lang=cn&!document=1#MediaInputInfo) 的 Id。</li>
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件名称
	// <li>若流程由 [ProcessMedia](http://139.199.214.26/document/product/266/30248) 发起，该字段表示 [MediaInfo](http://139.199.214.26/document/product/266/16694?!preview&preview_docmenu=1&lang=cn&!document=1#MediaInfo) 的 BasicInfo.Name；</li>
	// <li>若流程由 [ProcessMediaByUrl](http://139.199.214.26/document/product/266/31058?!document=1&!preview) 发起，该字段表示 [MediaInputInfo](http://139.199.214.26/document/product/266/16694?!preview&preview_docmenu=1&lang=cn&!document=1#MediaInputInfo) 的 Name。</li>
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 媒体文件地址
	// <li>若流程由 [ProcessMedia](http://139.199.214.26/document/product/266/30248) 发起，该字段表示 [MediaInfo](http://139.199.214.26/document/product/266/16694?!preview&preview_docmenu=1&lang=cn&!document=1#MediaInfo) 的 BasicInfo.MediaUrl；</li>
	// <li>若流程由 [ProcessMediaByUrl](http://139.199.214.26/document/product/266/31058?!document=1&!preview) 发起，该字段表示 [MediaInputInfo](http://139.199.214.26/document/product/266/16694?!preview&preview_docmenu=1&lang=cn&!document=1#MediaInputInfo) 的 Url。</li>
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 原始视频的元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 视频处理任务的执行状态与结果。
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitempty" name:"MediaProcessResultSet" list`

	// 视频内容审核任务的执行状态与结果。
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitempty" name:"AiContentReviewResultSet" list`

	// 视频内容分析任务的执行状态与结果。
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitempty" name:"AiAnalysisResultSet" list`

	// 任务流的优先级，取值范围为 [-10, 10]。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式。
	// <li>Finish：只有当任务流全部执行完毕时，才发起一次事件通知；</li>
	// <li>Change：只要任务流中每个子任务的状态发生变化，都进行事件通知；</li>
	// <li>None：不接受该任务流回调。</li>
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`
}

type ProcessMediaByUrlRequest struct {
	*tchttp.BaseRequest

	// 输入视频信息，包括视频 URL ， 名称、视频自定义 ID。
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// 输出文件 COS 路径信息。
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 250 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaByUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaByUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 250 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullEventsRequest struct {
	*tchttp.BaseRequest

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PullEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件列表。
		EventSet []*EventContent `json:"EventSet,omitempty" name:"EventSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullFileTask struct {

	// 转拉上传任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转拉上传完成后生成的视频 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 转拉上传完成后生成的播放地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 若转拉上传时指定了视频处理流程，则该参数为流程任务 ID。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`
}

type SampleSnapshotTaskInput struct {

	// 采样截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type SearchMediaRequest struct {
	*tchttp.BaseRequest

	// 搜索文本，模糊匹配媒体文件名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：64 个字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 标签集合，匹配集合中任意元素。
	// <li>单个标签长度限制：8 个字符。</li>
	// <li>数组长度限制：10。</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 分类 ID 集合，匹配集合指定 ID 的分类及其所有子类。数组长度限制：10。
	ClassIds []*int64 `json:"ClassIds,omitempty" name:"ClassIds" list`

	// 创建时间的开始时间。
	// <li>大于等于开始时间。</li>
	// <li>格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 创建时间的结束时间。
	// <li>小于结束时间。</li>
	// <li>格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 媒体文件来源，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 推流[直播码](https://cloud.tencent.com/document/product/267/5959)。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 直播录制文件的唯一标识。
	Vid *string `json:"Vid,omitempty" name:"Vid"`

	// 排序方式。
	// <li>Sort.Field 可选值：CreateTime</li>
	// <li>指定 Text 搜索时，将根据匹配度排序，该字段无效</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 偏移量。
	// <li>默认值：0。</li>
	// <li>取值范围：Offset + Limit 不超过 5000。</li>
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *SearchMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数
	// <li>最大值：5000，即，当命中记录数超过 5000，该字段将返回 5000，而非实际命中总数。</li>
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 媒体文件信息列表，只包含基础信息（BasicInfo）
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimpleHlsClipRequest struct {
	*tchttp.BaseRequest

	// 需要裁剪的腾讯云点播 HLS 视频 URL。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 裁剪的开始偏移时间，单位秒。默认 0，即从视频开头开始裁剪。负数表示距离视频结束多少秒开始裁剪。比如 -10 表示从倒数第 10 秒开始裁剪。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 裁剪的结束偏移时间，单位秒。默认 0，即裁剪到视频尾部。负数表示距离视频结束多少秒结束裁剪。比如 -10 表示到倒数第 10 秒结束裁剪。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

func (r *SimpleHlsClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimpleHlsClipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimpleHlsClipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 裁剪后的视频地址。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 裁剪后的视频元信息。目前`Size`，`Rotate`，`VideoDuration`，`AudioDuration` 几个字段暂时缺省，没有真实数据。
		MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SimpleHlsClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimpleHlsClipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SnapshotByTimeOffset2017 struct {

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 截图的具体时间点，单位：毫秒。
	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 截图输出文件地址。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type SnapshotByTimeOffsetTask2017 struct {

	// 截图任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 截图文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/11702#.E6.8C.87.E5.AE.9A.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 截图结果信息。
	SnapshotInfoSet []*SnapshotByTimeOffset2017 `json:"SnapshotInfoSet,omitempty" name:"SnapshotInfoSet" list`
}

type SnapshotByTimeOffsetTaskInput struct {

	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截图时间点列表，单位为秒。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet" list`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type SortBy struct {

	// 排序字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序）
	Order *string `json:"Order,omitempty" name:"Order"`
}

type TaskSimpleInfo struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型，取值：
	// <li>Procedure：视频处理任务；</li>
	// <li>EditMedia：视频编辑任务</li>
	// <li>WechatDistribute：微信发布任务。</li>
	// 兼容 2017 版的任务类型：
	// <li>Transcode：视频转码任务；</li>
	// <li>SnapshotByTimeOffset：视频截图任务；</li>
	// <li>Concat：视频拼接任务；</li>
	// <li>Clip：视频剪辑任务；</li>
	// <li>ImageSprites：截取雪碧图任务。</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 任务开始执行时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。若任务尚未开始，该字段为空。
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// 任务结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。若任务尚未完成，该字段为空。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type TempCertificate struct {

	// 临时安全证书 Id。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 临时安全证书 Key。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Token 值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 证书无效的时间，返回 Unix 时间戳，精确到秒。
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type TranscodePlayInfo2017 struct {

	// 播放地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/11701#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type TranscodeTask2017 struct {

	// 转码任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 被转码文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 被转码文件名称。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 视频时长，单位：秒。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 封面地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 视频转码后生成的播放信息。
	PlayInfoSet []*TranscodePlayInfo2017 `json:"PlayInfoSet,omitempty" name:"PlayInfoSet" list`
}

type TranscodeTaskInput struct {

	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type WatermarkInput struct {

	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`
}

type WechatPublishTask struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// WAITING：等待中；
	// PROCESSING：处理中；
	// FINISH：已完成。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 发布视频文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 微信发布模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 发布视频所对应的转码模板 ID，为 0 代表原始视频。
	SourceDefinition *uint64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// 微信发布状态，取值：
	// <li>FAIL：失败；</li>
	// <li>SUCCESS：成功；</li>
	// <li>AUDITNOTPASS：审核未通过；</li>
	// <li>NOTTRIGGERED：尚未发起微信发布。</li>
	WechatStatus *string `json:"WechatStatus,omitempty" name:"WechatStatus"`

	// 微信 Vid。
	WechatVid *string `json:"WechatVid,omitempty" name:"WechatVid"`

	// 微信地址。
	WechatUrl *string `json:"WechatUrl,omitempty" name:"WechatUrl"`
}
