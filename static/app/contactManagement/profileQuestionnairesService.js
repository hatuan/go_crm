/**
 * Created by tuanha-01 on 5/30/2016.
 */
define(['angularAMD', 'ajaxService'], function(angularAMD) {
    var injectParams = ['$interpolate', 'ajaxService'];

    profileQuestionnairesService = function($interpolate, ajaxService) {

        this.getProfileQuestionnaires = function(data, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(data, "/api/profilequestionnaires", successFunction, errorFunction);
        };

        this.updateProfileQuestionnaire = function(data, successFunction, errorFunction) {
            ajaxService.AjaxPost(data, "/api/profilequestionnaires", successFunction, errorFunction);
        };

        this.deleteProfileQuestionnaire = function(data, successFunction, errorFunction) {
            ajaxService.AjaxDelete(data, "/api/profilequestionnaires", successFunction, errorFunction);
        };

        this.getProfileQuestionnaire = function(data, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(data, "/api/profilequestionnaire", successFunction, errorFunction);
        };

        this.getProfileQuestionnaireLinesAndRatings = function(dataUrl, successFunction, errorFunction) {
            var getPath = '/api/profilequestionnaire/{{HeaderID}}';
            getPath = $interpolate(getPath)(dataUrl);
            ajaxService.AjaxGet(getPath, successFunction, errorFunction);
        };

        this.updateProfileQuestionnaireLinesAndRatings = function(dataUrl, data, successFunction, errorFunction) {
            var getPath = '/api/profilequestionnaire/{{HeaderID}}';
            getPath = $interpolate(getPath)(dataUrl);

            ajaxService.AjaxPost(data, getPath, successFunction, errorFunction);
        };
    };

    profileQuestionnairesService.$inject = injectParams;
    angularAMD.service('profileQuestionnairesService', profileQuestionnairesService);
});