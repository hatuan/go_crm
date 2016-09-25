define(['application-configuration', 'ajaxService'], function (app) {

    app.register.service('sqlParseService', ['ajaxService', function (ajaxService) {
        this.getSqlCondition = function (sqlCondition, successFunction, errorFunction) {
            ajaxService.AjaxPost(sqlCondition, "/api/sqlparse", successFunction, errorFunction);
        };
    }]);

});