myTodoList.controller('todoController', ['$scope', 'todoFactory', function ($scope, todoFactory) {

    $scope.todos;
    $scope.status;

    getAllTodos();

    function getAllTodos() {
        todoFactory.getAllTodos()
            .success(function (todos) {
                $scope.todos = todos;
            })
            .error(function (error) {
                $scope.status = 'Unable to load customer data: ' + error.message;
            });

    };

}]);