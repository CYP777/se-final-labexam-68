package test

import (
	"se-lab-exam-final/entity"
	"se-lab-exam-final/service"
	"testing"
	"time"

	"github.com/onsi/gomega"
)

// Test 1: กรณีข้อมูลถูกต้องครบถ้วน (Success Case)
func TestProjectProposal_Success(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// สร้างข้อมูลจำลองที่ถูกต้องตามเงื่อนไขทุกอย่าง
	proposal := entity.ProjectProposal{
		Projectname:    "Software Engineering Project", // 5-100 chars
		Description:    "This is a valid description.", // max 500
		StartDate:      time.Now(),
		DueDate:        time.Now().Add(24 * time.Hour), // DueDate > StartDate
		Status:         "pending",                      // oneof: pending, approved, rejected
		RequstedAmount: 5000.00,                        // 1 - 999999
		CoverPath:      "uploads/cover.jpg",            // required
	}

	// เรียกใช้ฟังก์ชัน Validate
	ok, err := service.ValidateProjectProposal(&proposal)

	// ตรวจสอบผลลัพธ์
	g.Expect(ok).To(gomega.BeTrue()) // ต้องได้ True
	g.Expect(err).To(gomega.BeNil()) // ต้องไม่มี Error
}

// Test 2: ตรวจสอบ ProjectName (Required)
func TestProjectProposal_ProjectNameRequired(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	proposal := entity.ProjectProposal{
		Projectname:    "", // ผิด: เป็นค่าว่าง
		Description:    "Valid description",
		StartDate:      time.Now(),
		DueDate:        time.Now().Add(24 * time.Hour),
		Status:         "pending",
		RequstedAmount: 5000.00,
		CoverPath:      "uploads/cover.jpg",
	}

	ok, err := service.ValidateProjectProposal(&proposal)

	g.Expect(ok).To(gomega.BeFalse()) // ต้องได้ False
	g.Expect(err).ToNot(gomega.BeNil()) // ต้องมี Error
	g.Expect(err.Error()).To(gomega.ContainSubstring("Project name is required")) // ข้อความ Error ต้องตรงกับ Tag
}

// Test 3: ตรวจสอบ ProjectName (Min 5 ตัวอักษร)
func TestProjectProposal_ProjectNameMinLength(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	proposal := entity.ProjectProposal{
		Projectname:    "ABC", // ผิด: สั้นกว่า 5 ตัวอักษร
		Description:    "Valid description",
		StartDate:      time.Now(),
		DueDate:        time.Now().Add(24 * time.Hour),
		Status:         "pending",
		RequstedAmount: 5000.00,
		CoverPath:      "uploads/cover.jpg",
	}

	ok, err := service.ValidateProjectProposal(&proposal)

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.ContainSubstring("Project name must be between 5 to 100 characters"))
}

// Test 4: ตรวจสอบ DueDate (ต้องมากกว่า StartDate)
func TestProjectProposal_DueDateBeforeStartDate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	proposal := entity.ProjectProposal{
		Projectname:    "Valid Project",
		Description:    "Valid description",
		StartDate:      time.Now(),
		DueDate:        time.Now().Add(-24 * time.Hour), // ผิด: วันกำหนดส่งอยู่ก่อนวันเริ่ม (ย้อนหลัง)
		Status:         "pending",
		RequstedAmount: 5000.00,
		CoverPath:      "uploads/cover.jpg",
	}

	ok, err := service.ValidateProjectProposal(&proposal)

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.ContainSubstring("Due date must be after start date"))
}

// Test 5: ตรวจสอบ Status (ต้องเป็น pending, approved, rejected)
func TestProjectProposal_StatusInvalid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	proposal := entity.ProjectProposal{
		Projectname:    "Valid Project",
		Description:    "Valid description",
		StartDate:      time.Now(),
		DueDate:        time.Now().Add(24 * time.Hour),
		Status:         "unknown", // ผิด: ไม่ได้อยู่ในลิสต์ที่กำหนด
		RequstedAmount: 5000.00,
		CoverPath:      "uploads/cover.jpg",
	}

	ok, err := service.ValidateProjectProposal(&proposal)

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.ContainSubstring("Status must be either pending or approved or rejected"))
}

// Test 6: ตรวจสอบ RequestedAmount (Range 1-999999)
func TestProjectProposal_RequestedAmountRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// กรณีเกิน 999999
	proposal := entity.ProjectProposal{
		Projectname:    "Valid Project",
		Description:    "Valid description",
		StartDate:      time.Now(),
		DueDate:        time.Now().Add(24 * time.Hour),
		Status:         "pending",
		RequstedAmount: 1000000.00, // ผิด: เกิน 999,999
		CoverPath:      "uploads/cover.jpg",
	}

	ok, err := service.ValidateProjectProposal(&proposal)

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.ContainSubstring("Requested amount must be between 1 and 999999"))
}