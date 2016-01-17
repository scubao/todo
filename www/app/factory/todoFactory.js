myTodoList.factory('entryFactory',['$http', function($http) {

    var urlBase = "http://127.0.0.1:8080/todo";
    var entryFactory = {};
    
    entryFactory.getEntries = function () {
        return $http.get(urlBase);
    };

    entryFactory.getEntry = function (id) {
        return $http.get(urlBase+'/'+id);
    };

    entryFactory.createEntry = function (name) {
        var todoentry = {name: name};
        return $http.post(urlBase, todoentry);
    };

    entryFactory.updateEntry = function (entry) {
        return $http.put(urlBase + '/' + entry.id, entry)
    };

    entryFactory.deleteEntry = function (id) {
        return $http.delete(urlBase + '/' + id);
    };
    
    return entryFactory;
    
}]);