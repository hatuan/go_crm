/**
 * Created by tuanha-01 on 5/30/2016.
 */
define(['angularAMD', 'ajaxService'], function (angularAMD) {
    var injectParams = ['$interpolate', 'ajaxService'];

    profileQuestionnairesService = function ($interpolate, ajaxService) {
        
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

        this.getProfileQuestionnaireLines = function (dataUrl, successFunction, errorFunction) {
            var getPath = '/api/profilequestionnaire/{{HeaderID}}/lines';
            getPath = $interpolate(getPath)(dataUrl);
            ajaxService.AjaxGet(getPath, successFunction, errorFunction);
        };

        this.updateProfileQuestionnaireLines = function (dataUrl, data, successFunction, errorFunction) {
            var getPath = '/api/profilequestionnaire/{{HeaderID}}/lines';
            getPath = $interpolate(getPath)(dataUrl);

            ajaxService.AjaxPost(data, getPath, successFunction, errorFunction);
        };
    };

    profileQuestionnairesService.$inject = injectParams;
    angularAMD.service('profileQuestionnairesService', profileQuestionnairesService);
/*
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

        this.getProfileQuestionnaireLines = function (data, successFunction, errorFunction) {
            var getPath = '/api/profilequestionnaire/{{HeaderID}}/lines';
            getPath = $interpolate(getPath)(data);
            ajaxService.AjaxGet(getPath, successFunction, errorFunction);
        };
    }]);
*/
});
