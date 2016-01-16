myTodoList.factory('todoFactory',['$http', function($http) {
 
    console.log("todoFactory used.");

    var urlBase = "http://localhost:8080/todo";
    var todoFactory = {};
    
    todoFactory.getAllTodos = function () {
        return $http.get(urlBase);
    };
    
    return todoFactory;
    
}]);