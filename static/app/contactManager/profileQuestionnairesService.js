/**
 * Created by tuanha-01 on 5/30/2016.
 */
define(['angularAMD', 'ajaxService'], function (angularAMD) {

    angularAMD.service('profileQuestionnairesService', ['ajaxService', function (ajaxService) {
        
        this.getProfileQuestionnaires = function (data, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(data, "/api/profilequestionnaires", successFunction, errorFunction);
        };

        this.updateProfileQuestionnaire = function (data, successFunction, errorFunction) {
            ajaxService.AjaxPost(data, "/api/profilequestionnaires", successFunction, errorFunction);
        };

        this.deleteProfileQuestionnaire = function (data, successFunction, errorFunction) {
            ajaxService.AjaxDelete(data, "/api/profilequestionnaires", successFunction, errorFunction);
        };

        this.getProfileQuestionnaire = function (data, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(data, "/api/profilequestionnaire", successFunction, errorFunction);
        };    
    }]);
});
