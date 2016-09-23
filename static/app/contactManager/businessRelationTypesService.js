/**
 * Created by tuanha-01 on 5/30/2016.
 */
define(['application-configuration', 'ajaxService'], function (app) {

    app.register.service('businessRelationTypesService', ['ajaxService', function (ajaxService) {
        
        this.getBusinessRelationTypes = function (businessRelationType, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(businessRelationType, "/api/businessrelationtypes", successFunction, errorFunction);
        };

        this.updateBusinessRelationType = function (businessRelationType, successFunction, errorFunction) {
            ajaxService.AjaxPost(businessRelationType, "/api/businessrelationtypes", successFunction, errorFunction);
        };

        this.deleteBusinessRelationType = function (data, successFunction, errorFunction) {
            ajaxService.AjaxDelete(data, "/api/businessrelationtypes", successFunction, errorFunction);
        };

        this.getBusinessRelationType = function (businessRelationType, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(businessRelationType, "/api/businessrelationtype", successFunction, errorFunction);
        };    
    }]);
});
