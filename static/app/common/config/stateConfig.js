/**
 * Created by tuanha-01 on 5/11/2016.
 */
define(['angularAMD'], function (angularAMD) {

    var injectParams = ['$stateProvider', '$urlRouterProvider', '$authProvider'];
    stateConfig = function ($stateProvider, $urlRouterProvider, $authProvider) {

        $urlRouterProvider.otherwise("/");

        $stateProvider
            .state('login', {
                url: '/login',
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/user/login.html',
                        controller: 'LoginController',
                        controllerUrl: 'app/user/loginController',
                    })
                },
                resolve: {
                    skipIfAuthenticated: _skipIfAuthenticated
                }
            })
            .state('home', {
                url: '/',
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/main/home.html',
                        controller: 'HomeController',
                        controllerUrl: 'app/main/homeController',
                    })
                },
                resolve: {
                    skipIfAuthenticated: _skipIfAuthenticated
                }
            })
            .state('preference', {
                url: '/preference',
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/user/preference.html',
                        controller: 'PreferenceController',
                        controllerUrl: 'app/user/preferenceController',
                    })
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('module', {
                abstract: true,
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/module/module.html',
                        controller: 'ModuleController',
                        controllerUrl: 'app/module/moduleController'
                    })
                },
            })
            .state('module.master', {
                url: '/module/master',
                views: {
                    "master@module": angularAMD.route({
                        templateUrl: 'app/module/master.html',
                    })
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('module.contactManagement', {
                url: '/module/contactmanagement',
                views: {
                    "master@module": angularAMD.route({
                        templateUrl: 'app/module/contactManagement.html',
                    })
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('setup', {
                url: "/setup/:setupModule"
            })
            .state('businessRelationType', {
                url: "/businessRelationType",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/businessRelationTypes.html',
                        controller: 'BusinessRelationTypesController',
                        controllerUrl: 'app/contactManagement/businessRelationTypesController',
                    })
                },
                params: {
                    businessRelationTypeID: null
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('businessRelationTypeMaintenance', {
                url: "/businessRelationTypeMaintenance",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/businessRelationTypeMaintenance.html',
                        controller: 'BusinessRelationTypeMaintenanceController',
                        controllerUrl: 'app/contactManagement/businessRelationTypeMaintenanceController',
                    })
                },
                params: {
                    id: null
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('businessRelationSector', {
                url: "/businessrelationsector",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/businessRelationSectors.html',
                        controller: 'BusinessRelationSectorsController',
                        controllerUrl: 'app/contactManagement/businessRelationSectorsController',
                    })
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('businessRelationSectorMaintenance', {
                url: "/businessRelationSectorMaintenance",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/businessRelationSectorMaintenance.html',
                        controller: 'BusinessRelationSectorMaintenanceController',
                        controllerUrl: 'app/contactManagement/businessRelationSectorMaintenanceController',
                    })
                },
                params: {
                    id: null
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('profileQuestionnaire', {
                url: "/profilequestionnaire",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/profileQuestionnaires.html',
                        controller: 'ProfileQuestionnairesController',
                        controllerUrl: 'app/contactManagement/profileQuestionnairesController',
                    })
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('profileQuestionnaireMaintenance', {
                url: "/profileQuestionnaireMaintenance",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/profileQuestionnaireMaintenance.html',
                        controller: 'ProfileQuestionnaireMaintenanceController',
                        controllerUrl: 'app/contactManagement/profileQuestionnaireMaintenanceController',
                    })
                },
                params: {
                    ID: null
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            .state('profileQuestionnaireLinesMaintenance', {
                url: "/profileQuestionnaireLinesMaintenance",
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/contactManagement/profileQuestionnaireLinesMaintenance.html',
                        controller: 'ProfileQuestionnaireLinesMaintenanceController',
                        controllerUrl: 'app/contactManagement/profileQuestionnaireLinesMaintenanceController',
                    })
                },
                params: {
                    headerID: null
                },
                resolve: {
                    redirectIfNotAuthenticated: _redirectIfNotAuthenticated,
                }
            })
            ;

        function _skipIfAuthenticated($q, $state, $auth) {
            var defer = $q.defer();
            if ($auth.isAuthenticated()) {
                defer.resolve(); // always return defer.resolve()
            } else {
                defer.resolve(); // always return defer.resolve()
            }
            return defer.promise;
        }

        function _redirectIfNotAuthenticated($q, $state, $auth) {
            var defer = $q.defer();
            if ($auth.isAuthenticated()) {
                defer.resolve(); // always return defer.resolve()
            } else {
                setTimeout(function () {
                    $state.go('login');
                }, 10);
                defer.resolve(); // always return defer.resolve()
            }
            return defer.promise;
        }
    }

    stateConfig.$inject = injectParams;

    return stateConfig;
});