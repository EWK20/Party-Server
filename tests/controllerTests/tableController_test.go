package controller_test

import (
	"testing"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockTableService struct {
	tableMock mock.Mock
}

func (s *MockTableService) FindAll() ([]dto.TableResDto, error) {
	args := s.tableMock.Called()
	if args.Error(1) != nil {
		return []dto.TableResDto{}, args.Error(1)
	}
	return args.Get(0).([]dto.TableResDto), nil
}

func (s *MockTableService) FindById(id int) (dto.TableResDto, error) {
	args := s.tableMock.Called(id)
	if args.Error(1) != nil {
		return dto.TableResDto{}, args.Error(1)
	}
	return args.Get(0).(dto.TableResDto), nil
}

func (s *MockTableService) Save(req dto.TableReqDto) (dto.TableResDto, error) {
	args := s.tableMock.Called(req)
	if args.Error(1) != nil {
		return dto.TableResDto{}, args.Error(1)
	}
	return args.Get(0).(dto.TableResDto), nil
}

func (s *MockTableService) CheckSpace() (int, bool) {
	args := s.tableMock.Called()
	if args.Bool(1) == false {
		return 0, args.Bool(1)
	}
	return args.Int(0), true
}

// var (
// 	tableRepository repository.TableRepository = repository.NewTableRepository()

// 	tableService service.TableService = service.NewTableService(tableRepository)

// 	tableController controller.TableController = controller.NewTableController(tableService)
// )

func TestGetTables(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTableRes := []dto.TableResDto{
			{
				Id:       1,
				Capacity: 10,
			},
			{
				Id:       2,
				Capacity: 2,
			},
			{
				Id:       3,
				Capacity: 2,
			},
			{
				Id:       4,
				Capacity: 5,
			},
			{
				Id:       5,
				Capacity: 15,
			},
		}

		mockTableService := new(MockTableService)

		mockTableService.tableMock.On("FindAll", mock.AnythingOfType("*gin.Context")).Return(mockTableRes, nil)

		// A response recorder for getting written HTTP response
		// rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("tables", mockTableRes)
		})
	})

}
