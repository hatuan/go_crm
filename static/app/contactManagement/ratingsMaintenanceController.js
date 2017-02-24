/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function(angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants', 'profileQuestionnaireLineEditRatings'];

    var ratingsMaintenanceController = function($scope, $rootScope, $state, $window, moment, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, Constants, profileQuestionnaireLineEditRatings) {

        $scope.ProfileQuestionnaireLineEditRatings = profileQuestionnaireLineEditRatings;
        if ($scope.ProfileQuestionnaireLineEditRatings.Ratings === null)
            $scope.ProfileQuestionnaireLineEditRatings.Ratings = [];

        $scope.ProfileQuestionnaireLines = [];
        $scope.ProfileQuestionnaires = [];
        $scope.Constants = Constants;

        $scope.initializeController = function() {
            $scope.selectProfileQuestionnaireHeaderID = "";

            $scope.getProfileQuestionnaires();
        };

        $scope.getProfileQuestionnaires = function() {
            var profileQuestionnaireInquiry = new Object();

            profileQuestionnaireInquiry.Search = "";
            profileQuestionnaireInquiry.SortExpression = "Code";
            profileQuestionnaireInquiry.SortDirection = "ASC";
            profileQuestionnaireInquiry.FetchSize = ""; //get all

            profileQuestionnairesService.getProfileQuestionnaires(profileQuestionnaireInquiry, $scope.profileQuestionnairesInquiryCompleted, $scope.profileQuestionnairesInquiryError);
        };

        $scope.profileQuestionnairesInquiryCompleted = function(response, status) {
            $scope.ProfileQuestionnaires = response.Data.ProfileQuestionnaires;
        };

        $scope.profileQuestionnairesInquiryError = function(response, status) {
            alertsService.RenderErrorMessage(response.Error);
        };

        $scope.selectProfileQuestionnaire = function(_profileQuestionnaireHeaderID) {
            $scope.selectProfileQuestionnaireHeaderID = _profileQuestionnaireHeaderID;

            $scope.getProfileQuestionnaireLines(_profileQuestionnaireHeaderID);
        };

        var saveCurrentPointsOfProfileQuestionnaireLinesToRatings = function() {
            for (var i = 0, len = $scope.ProfileQuestionnaireLines.length; i < len; i++) {
                var points = parseInt($scope.ProfileQuestionnaireLines[i].Points);
                if (!isNaN(points)) {
                    var rating = $scope.ProfileQuestionnaireLineEditRatings.Ratings.find(function(el) {
                        return el.RatingProfileQuestionnaireLineID == $scope.ProfileQuestionnaireLines[i].ID
                    })

                    if (rating) { //replace point
                        rating.Points = points;
                    } else { //append to ratings
                        var newRating = new Object();
                        newRating.ProfileQuestionnaireHeaderID = $scope.ProfileQuestionnaireLineEditRatings.ProfileQuestionnaireHeaderID;
                        newRating.ProfileQuestionnaireLineID = $scope.ProfileQuestionnaireLineEditRatings.ID;
                        newRating.RatingProfileQuestionnaireHeaderID = $scope.ProfileQuestionnaireLines[i].ProfileQuestionnaireHeaderID;
                        newRating.RatingProfileQuestionnaireLineID = $scope.ProfileQuestionnaireLines[i].ID;
                        newRating.Points = points;

                        $scope.ProfileQuestionnaireLineEditRatings.Ratings.push(newRating);
                    }
                } else { //clear current Ratings point
                    var rating = $scope.ProfileQuestionnaireLineEditRatings.Ratings.find(function(el) {
                        return el.RatingProfileQuestionnaireLineID == $scope.ProfileQuestionnaireLines[i].ID
                    })
                    if (rating) {
                        rating.Points = 0;
                    }
                }
            }
        };

        $scope.getProfileQuestionnaireLines = function(_profileQuestionnaireHeaderID) {
            if (!angular.element("[name='ProfileQuestionnaireLinesMaintenanceForm']").controller("form").validate())
                return;

            //save current points of ProfileQuestionnaireLines to ratings
            saveCurrentPointsOfProfileQuestionnaireLinesToRatings();

            //display new ProfileQuestionnaireLines
            profileQuestionnairesService.getProfileQuestionnaireLines({ HeaderID: _profileQuestionnaireHeaderID }, $scope.getProfileQuestionnaireLinesCompleted, $scope.getProfileQuestionnaireLinesError);
        };

        $scope.getProfileQuestionnaireLinesCompleted = function(response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);

            var profileQuestionnaireLines = response.Data.ProfileQuestionnaireLines;
            var profileQuestionnaireLineDeletes = [];
            var found_question = false;
            for (var i = 0, len = profileQuestionnaireLines.length; i < len; i++) {
                profileQuestionnaireLines[i].RecCreated = new moment.unix(profileQuestionnaireLines[i].RecCreated).toDate();
                profileQuestionnaireLines[i].RecModified = new moment.unix(profileQuestionnaireLines[i].RecModified).toDate();

                if (profileQuestionnaireLines[i].ProfileQuestionnaireHeaderID == $scope.ProfileQuestionnaireLineEditRatings.ProfileQuestionnaireHeaderID) {
                    if (profileQuestionnaireLines[i].LineNo == $scope.ProfileQuestionnaireLineEditRatings.LineNo) {
                        found_question = true;
                        profileQuestionnaireLineDeletes.push(profileQuestionnaireLines[i])
                    } else if (found_question && profileQuestionnaireLines[i].Type == Constants.ProfileQuestionaireLineTypes[1].Code) {
                        profileQuestionnaireLineDeletes.push(profileQuestionnaireLines[i])
                    } else if (found_question && profileQuestionnaireLines[i].Type == Constants.ProfileQuestionaireLineTypes[0].Code) {
                        found_question = false;
                    }
                }

                var ratingPoints = $scope.ProfileQuestionnaireLineEditRatings.Ratings.find(function(el) {
                    return el.RatingProfileQuestionnaireLineID == profileQuestionnaireLines[i].ID
                })
                if (ratingPoints) {
                    profileQuestionnaireLines[i].Points = ratingPoints.Points;
                }
            }

            profileQuestionnaireLines = profileQuestionnaireLines.filter(function(el, index, array) {
                return !profileQuestionnaireLineDeletes.find(function(_el) {
                    return _el.ID === el.ID
                })
            })
            $scope.ProfileQuestionnaireLines = profileQuestionnaireLines;
        };

        $scope.getProfileQuestionnaireLinesError = function(response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.validationOptions = {
            rules: {
                "Points[]": {
                    number: true
                }
            }
        };

        $scope.ok = function(form) {
            if (form.validate()) {
                //save current points of ProfileQuestionnaireLines to ratings
                saveCurrentPointsOfProfileQuestionnaireLinesToRatings();
                var profileQuestionnaireLineEditRatings = $scope.ProfileQuestionnaireLineEditRatings;
                profileQuestionnaireLineEditRatings.RecCreatedByUser = $scope.RecCreatedByUser;
                profileQuestionnaireLineEditRatings.RecCreated = new moment($scope.RecCreated).unix();
                profileQuestionnaireLineEditRatings.RecModifiedByID = $rootScope.currentUser.ID;
                profileQuestionnaireLineEditRatings.RecModifiedByUser = $rootScope.currentUser.Name;
                profileQuestionnaireLineEditRatings.RecModified = new moment($scope.RecModified).unix();

                profileQuestionnairesService.updateProfileQuestionnaireLineRatings({
                        ProfileQuestionnaireLine: profileQuestionnaireLineEditRatings
                    },
                    $scope.updateProfileQuestionnaireLineRatingsCompleted,
                    $scope.updateProfileQuestionnaireLineRatingsError);
            }
        };

        $scope.updateProfileQuestionnaireLineRatingsCompleted = function(response, status) {
            $scope.ProfileQuestionnaireLineEditRatings.Ratings = response.Data.Ratings;

            $uibModalInstance.close($scope.ProfileQuestionnaireLineEditRatings);
        };

        $scope.updateProfileQuestionnaireLineRatingsError = function(response, status) {

        };

        $scope.cancel = function(form) {
            $uibModalInstance.dismiss('cancel');
        };

        $scope.changeProfileQuestionaire = function() {

        };
    };

    ratingsMaintenanceController.$inject = injectParams;
    angularAMD.controller('RatingsMaintenanceController', ratingsMaintenanceController);
});