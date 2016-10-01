define(['angularAMD', 'ajaxService'], function (angularAMD) {

    angularAMD.service('sqlParseService', ['ajaxService', function (ajaxService) {
        this.getSqlCondition = function (sqlCondition, successFunction, errorFunction) {
            ajaxService.AjaxPost(sqlCondition, "/api/sqlparse", successFunction, errorFunction);
        };
    }]);

});