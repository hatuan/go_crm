/**
 * Created by tuanha-01 on 5/23/2016.
 */
define(['application-configuration', 'sqlParseService'], function (app) {

    var injectParams = ['sqlParseService'];

    var searchFn = function (sqlParseService) {
        return {
            restrict: 'EA',
            templateUrl: 'app/common/directives/search.html',
            scope: {
                searchConditionObjects: '@searchConditionObjects',
                parentSearch: '&search'
            },
            link: function (scope, element, attrs, modelCtrl) {
                
            },
            controller: function($scope){
                $scope.initializeController = function () {
                    $scope.searchSqlCondition = "";

                    $scope.searchConditions = [];
                    $scope.addSearchCondition();
                }

                $scope.startSearch = function() {
                    var searchConditions = [];
                    
                    for(var _i = 0; _i < $scope.searchConditions.length; _i ++) {
                        searchConditions.push({
                            "ID": $scope.searchConditions[_i].Object.ID,
                            "Value": $scope.searchConditions[_i].Object.Value,
                        });
                    }

                    sqlParseService.getSqlCondition(searchConditions, $scope.searchCompleted, $scope.searchError);
                }

                $scope.clearSearch = function() {
                    $scope.searchConditions = [];
                    $scope.addSearchCondition();

                    $scope.searchSqlCondition = "";
                    $scope.parentSearch();
                }

                $scope.searchCompleted = function(response, status) {
                    var errs = response.Data.Errs;
                    var stmts = response.Data.Stmts;
                    
                    for(var _i = 0; _i < errs.length; _i ++) {
                        $scope.SearchConditions[_i].Error = "";
                        $scope.SearchConditions[_i].HasError = false;
                        $scope.SearchConditions[_i].Stmt = stmts[_i];
                        if (_i == 0)
                            $scope.searchSqlCondition = "(" + $scope.searchConditions[_i].Stmt +")";
                        else 
                            $scope.searchSqlCondition += " AND (" + $scope.searchConditions[_i].Stmt + ")";
                    }
                    $scope.searchSqlCondition = "(" + $scope.searchSqlCondition + ")";
                    $scope.parentSearch(searchSqlCondition);
                }

                $scope.searchError = function(response, status) {
                    var errs = response.Data.Errs;
                    var stmts = response.Data.Stmts;
                    for(var _i = 0; _i < errs.length; _i ++) {
                        $scope.searchConditions[_i].Error = errs[_i];
                        $scope.searchConditions[_i].HasError = errs[_i].length > 0;
                        $scope.searchConditions[_i].Stmt = stmts[_i];
                    }
                }

                 $scope.addSearchCondition = function(){
                    var searchCondition = new Object();
                    searchCondition.No = $scope.SearchConditions.length + 1;
                    searchCondition.Err = "";
                    searchCondition.HasError = false;
                    searchCondition.Stmt = "";
                    searchCondition.Objects = JSON.parse(JSON.stringify($scope.SearchConditionObjects));
                    searchCondition.Object = searchCondition.Objects[0];

                    $scope.SearchConditions.push(searchCondition);
                }
                
                $scope.removeSearchCondition = function(_index){
                    $scope.SearchConditions.splice(_index, 1);
                }

                $scope.initializeController();
            },
        }
    };

    searchFn.$inject = injectParams;

    app.register.directive('search', searchFn)
});
