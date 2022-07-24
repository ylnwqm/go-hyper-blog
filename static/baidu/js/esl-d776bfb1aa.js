(function(root) {
    if (root.esl && root.require && root.esl.version === root.require.version) {
        return
    }
    var define;
    var require;
    var esl;
    (function(global) {
        var modModules = {};
        var MODULE_PRE_DEFINED = 1;
        var MODULE_ANALYZED = 2;
        var MODULE_PREPARED = 3;
        var MODULE_DEFINED = 4;
        var modAutoDefineModules = {};

        function modFlagAutoDefine(id) {
            if (!modIs(id, MODULE_DEFINED)) {
                modAutoDefineModules[id] = 1
            }
        }
        var BUILDIN_MODULE = {
            require: globalRequire,
            exports: 1,
            module: 1
        };
        var actualGlobalRequire = createLocalRequire();
        var waitTimeout;
        var requireConf = {
            baseUrl: "./",
            paths: {},
            config: {},
            map: {},
            packages: [],
            shim: {},
            waitSeconds: 0,
            bundles: {},
            urlArgs: {}
        };

        function globalRequire(requireId, callback) {
            var invalidIds = [];

            function monitor(id) {
                if (id.indexOf(".") === 0) {
                    invalidIds.push(id)
                }
            }
            if (typeof requireId === "string") {
                monitor(requireId)
            } else {
                each(requireId, function(id) {
                    monitor(id)
                })
            }
            if (invalidIds.length > 0) {
                throw new Error("[REQUIRE_FATAL]Relative ID is not allowed in global require: " + invalidIds.join(", "))
            }
            var timeout = requireConf.waitSeconds;
            if (timeout && requireId instanceof Array) {
                if (waitTimeout) {
                    clearTimeout(waitTimeout)
                }
                waitTimeout = setTimeout(waitTimeoutNotice, timeout * 1e3)
            }
            return actualGlobalRequire(requireId, callback)
        }
        globalRequire.version = "2.1.4";
        globalRequire.loader = "esl";
        globalRequire.toUrl = actualGlobalRequire.toUrl;

        function waitTimeoutNotice() {
            var hangModules = [];
            var missModules = [];
            var hangModulesMap = {};
            var missModulesMap = {};
            var visited = {};

            function checkError(id, hard) {
                if (visited[id] || modIs(
                        id, MODULE_DEFINED)) {
                    return
                }
                visited[id] = 1;
                var mod = modModules[id];
                if (!mod) {
                    if (!missModulesMap[id]) {
                        missModulesMap[id] = 1;
                        missModules.push(id)
                    }
                } else if (hard || !modIs(id, MODULE_PREPARED) || mod.hang) {
                    if (!hangModulesMap[id]) {
                        hangModulesMap[id] = 1;
                        hangModules.push(id)
                    }
                    each(mod.depMs, function(dep) {
                        checkError(dep.absId, dep.hard)
                    })
                }
            }
            for (var id in modAutoDefineModules) {
                checkError(id, 1)
            }
            if (hangModules.length || missModules.length) {
                throw new Error("[MODULE_TIMEOUT]Hang(" + (hangModules.join(", ") || "none") + ") Miss(" + (missModules.join(", ") || "none") + ")")
            }
        }
        var wait4PreDefine = [];

        function modCompletePreDefine(currentId) {
            each(wait4PreDefine, function(mod) {
                modPreDefine(currentId, mod.deps, mod.factory)
            });
            wait4PreDefine.length = 0
        }

        function globalDefine(id, dependencies, factory) {
            if (factory == null) {
                if (dependencies == null) {
                    factory = id;
                    id = null
                } else {
                    factory = dependencies;
                    dependencies = null;
                    if (id instanceof Array) {
                        dependencies = id;
                        id = null
                    }
                }
            }
            if (factory == null) {
                return
            }
            var opera = window.opera;
            if (!id && document.attachEvent && !(
                    opera && opera.toString() === "[object Opera]")) {
                var currentScript = getCurrentScript();
                id = currentScript && currentScript.getAttribute("data-require-id")
            }
            if (id) {
                modPreDefine(id, dependencies, factory)
            } else {
                wait4PreDefine[0] = {
                    deps: dependencies,
                    factory: factory
                }
            }
        }
        globalDefine.amd = {};

        function moduleConfigGetter() {
            var conf = requireConf.config[this.id];
            if (conf && typeof conf === "object") {
                return conf
            }
            return {}
        }

        function modPreDefine(id, dependencies, factory) {
            if (!modModules[id]) {
                modModules[id] = {
                    id: id,
                    depsDec: dependencies,
                    deps: dependencies || ["require", "exports", "module"],
                    factoryDeps: [],
                    factory: factory,
                    exports: {},
                    config: moduleConfigGetter,
                    state: MODULE_PRE_DEFINED,
                    require: createLocalRequire(id),
                    depMs: [],
                    depMkv: {},
                    depRs: [],
                    hang: 0
                }
            }
        }

        function modPrepare(id) {
            var mod = modModules[id];
            if (!mod || modIs(id, MODULE_ANALYZED)) {
                return
            }
            var deps = mod.deps;
            var factory = mod.factory;
            var hardDependsCount = 0;
            if (typeof factory === "function") {
                hardDependsCount = Math.min(factory.length, deps.length);
                !mod.depsDec && factory.toString().replace(/(\/\*([\s\S]*?)\*\/|([^:]|^)\/\/(.*)$)/gm, "").replace(/require\(\s*(['"])([^'"]+)\1\s*\)/g, function($0, $1, depId) {
                    deps.push(depId)
                })
            }
            var requireModules = [];
            var depResources = [];
            each(deps, function(depId, index) {
                var idInfo = parseId(depId);
                var absId = normalize(idInfo.mod, id);
                var moduleInfo;
                var resInfo;
                if (absId && !BUILDIN_MODULE[absId]) {
                    if (idInfo.res) {
                        resInfo = {
                            id: depId,
                            mod: absId,
                            res: idInfo.res
                        };
                        depResources.push(depId);
                        mod.depRs.push(resInfo)
                    }
                    moduleInfo = mod.depMkv[absId];
                    if (!moduleInfo) {
                        moduleInfo = {
                            id: idInfo.mod,
                            absId: absId,
                            hard: index < hardDependsCount
                        };
                        mod.depMs.push(moduleInfo);
                        mod.depMkv[absId] = moduleInfo;
                        requireModules.push(absId)
                    }
                } else {
                    moduleInfo = {
                        absId: absId
                    }
                }
                if (index < hardDependsCount) {
                    mod.factoryDeps.push(resInfo || moduleInfo)
                }
            });
            mod.state = MODULE_ANALYZED;
            modInitFactoryInvoker(id);
            nativeAsyncRequire(requireModules);
            depResources.length && mod.require(depResources, function() {
                each(mod.depRs, function(res) {
                    if (!res.absId) {
                        res.absId = normalize(res.id, id)
                    }
                });
                modAutoDefine()
            })
        }

        function modAutoDefine() {
            for (var id in modAutoDefineModules) {
                modPrepare(id);
                modUpdatePreparedState(id);
                modTryInvokeFactory(id)
            }
        }

        function modUpdatePreparedState(id) {
            var updatingModules = {};
            update(id);

            function update(id) {
                modPrepare(id);
                if (!modIs(id, MODULE_ANALYZED)) {
                    return false
                }
                if (modIs(id, MODULE_PREPARED) || updatingModules[id]) {
                    return true
                }
                updatingModules[id] = 1;
                var mod = modModules[id];
                var prepared = true;
                each(mod.depMs, function(dep) {
                    prepared = update(
                        dep.absId) && prepared
                });
                prepared && each(mod.depRs, function(dep) {
                    prepared = !!dep.absId;
                    return prepared
                });
                if (prepared && !modIs(id, MODULE_PREPARED)) {
                    mod.state = MODULE_PREPARED
                }
                updatingModules[id] = 0;
                return prepared
            }
        }

        function modInitFactoryInvoker(id) {
            var mod = modModules[id];
            var invoking;
            mod.invokeFactory = invokeFactory;

            function invokeFactory() {
                if (invoking || mod.state !== MODULE_PREPARED) {
                    return
                }
                invoking = 1;
                var factoryReady = 1;
                each(mod.factoryDeps, function(dep) {
                    var depId = dep.absId;
                    if (!BUILDIN_MODULE[depId]) {
                        modTryInvokeFactory(depId);
                        return factoryReady = modIs(depId, MODULE_DEFINED)
                    }
                });
                if (factoryReady) {
                    try {
                        var factory = mod.factory;
                        var exports = typeof factory === "function" ? factory.apply(global, modGetModulesExports(mod.factoryDeps, {
                            require: mod.require,
                            exports: mod.exports,
                            module: mod
                        })) : factory;
                        if (exports != null) {
                            mod.exports = exports
                        }
                        mod.invokeFactory = null
                    } catch (ex) {
                        mod.hang = 1;
                        throw ex
                    }
                    modDefined(id)
                }
            }
        }

        function modIs(id, state) {
            return modModules[id] && modModules[id].state >= state
        }

        function modTryInvokeFactory(id) {
            var mod = modModules[id];
            if (mod && mod.invokeFactory) {
                mod.invokeFactory()
            }
        }

        function modGetModulesExports(modules, buildinModules) {
            var args = [];
            each(modules, function(id, index) {
                if (typeof id === "object") {
                    id = id.absId
                }
                args[index] = buildinModules[id] || modModules[id].exports
            });
            return args
        }
        var modDefinedListeners = {};

        function modAddDefinedListener(id, listener) {
            if (modIs(id, MODULE_DEFINED)) {
                listener();
                return
            }
            var listeners = modDefinedListeners[id];
            if (!listeners) {
                listeners = modDefinedListeners[id] = []
            }
            listeners.push(
                listener)
        }

        function modDefined(id) {
            var mod = modModules[id];
            mod.state = MODULE_DEFINED;
            delete modAutoDefineModules[id];
            var listeners = modDefinedListeners[id] || [];
            var len = listeners.length;
            while (len--) {
                listeners[len]()
            }
            listeners.length = 0;
            modDefinedListeners[id] = null
        }

        function nativeAsyncRequire(ids, callback, baseId) {
            var isCallbackCalled = 0;
            each(ids, function(id) {
                if (!(BUILDIN_MODULE[id] || modIs(id, MODULE_DEFINED))) {
                    modAddDefinedListener(id, tryFinishRequire);
                    (id.indexOf("!") > 0 ? loadResource : loadModule)(id,
                        baseId)
                }
            });
            tryFinishRequire();

            function tryFinishRequire() {
                if (typeof callback === "function" && !isCallbackCalled) {
                    var isAllCompleted = 1;
                    each(ids, function(id) {
                        if (!BUILDIN_MODULE[id]) {
                            return isAllCompleted = !!modIs(id, MODULE_DEFINED)
                        }
                    });
                    if (isAllCompleted) {
                        isCallbackCalled = 1;
                        callback.apply(global, modGetModulesExports(ids, BUILDIN_MODULE))
                    }
                }
            }
        }
        var loadingModules = {};

        function loadModule(moduleId) {
            if (loadingModules[moduleId] || modModules[moduleId]) {
                return
            }
            loadingModules[moduleId] = 1;
            var shimConf = requireConf.shim[moduleId];
            if (shimConf instanceof Array) {
                requireConf.shim[moduleId] = shimConf = {
                    deps: shimConf
                }
            }
            var shimDeps = shimConf && (shimConf.deps || []);
            if (shimDeps) {
                each(shimDeps, function(dep) {
                    if (!requireConf.shim[dep]) {
                        requireConf.shim[dep] = {}
                    }
                });
                actualGlobalRequire(shimDeps, load)
            } else {
                load()
            }

            function load() {
                var bundleModuleId = bundlesIndex[moduleId];
                createScript(bundleModuleId || moduleId, loaded)
            }

            function loaded() {
                if (shimConf) {
                    var exports;
                    if (typeof shimConf.init === "function") {
                        exports = shimConf.init.apply(global, modGetModulesExports(shimDeps, BUILDIN_MODULE))
                    }
                    if (exports == null && shimConf.exports) {
                        exports = global;
                        each(shimConf.exports.split("."), function(prop) {
                            exports = exports[prop];
                            return !!exports
                        })
                    }
                    globalDefine(moduleId, shimDeps, function() {
                        return exports || {}
                    })
                } else {
                    modCompletePreDefine(moduleId)
                }
                modAutoDefine()
            }
        }

        function loadResource(pluginAndResource, baseId) {
            if (modModules[pluginAndResource]) {
                return
            }
            var bundleModuleId = bundlesIndex[pluginAndResource];
            if (bundleModuleId) {
                loadModule(bundleModuleId);
                return
            }
            var idInfo = parseId(pluginAndResource);
            var resource = {
                id: pluginAndResource,
                state: MODULE_ANALYZED
            };
            modModules[pluginAndResource] = resource;

            function pluginOnload(value) {
                resource.exports = value || true;
                modDefined(pluginAndResource)
            }
            pluginOnload.fromText = function(id, text) {
                new Function(text)();
                modCompletePreDefine(id)
            };

            function load(plugin) {
                var pluginRequire = baseId ? modModules[baseId].require : actualGlobalRequire;
                plugin.load(idInfo.res, pluginRequire, pluginOnload,
                    moduleConfigGetter.call({
                        id: pluginAndResource
                    }))
            }
            load(actualGlobalRequire(idInfo.mod))
        }
        globalRequire.config = function(conf) {
            if (conf) {
                for (var key in requireConf) {
                    var newValue = conf[key];
                    var oldValue = requireConf[key];
                    if (!newValue) {
                        continue
                    }
                    if (key === "urlArgs" && typeof newValue === "string") {
                        requireConf.urlArgs["*"] = newValue
                    } else {
                        if (oldValue instanceof Array) {
                            oldValue.push.apply(oldValue, newValue)
                        } else if (typeof oldValue === "object") {
                            for (var k in newValue) {
                                oldValue[k] = newValue[k]
                            }
                        } else {
                            requireConf[key] = newValue
                        }
                    }
                }
                createConfIndex()
            }
        };
        createConfIndex();
        var pathsIndex;
        var packagesIndex;
        var mappingIdIndex;
        var bundlesIndex;
        var urlArgsIndex;

        function createKVSortedIndex(value, allowAsterisk) {
            var index = kv2List(value, 1, allowAsterisk);
            index.sort(descSorterByKOrName);
            return index
        }

        function createConfIndex() {
            requireConf.baseUrl = requireConf.baseUrl.replace(/\/$/, "") + "/";
            pathsIndex = createKVSortedIndex(requireConf.paths);
            mappingIdIndex = createKVSortedIndex(requireConf.map, 1);
            each(mappingIdIndex, function(item) {
                item.v = createKVSortedIndex(item.v)
            });
            var lastMapItem = mappingIdIndex[mappingIdIndex.length - 1];
            if (lastMapItem && lastMapItem.k === "*") {
                each(mappingIdIndex, function(item) {
                    if (item != lastMapItem) {
                        item.v = item.v.concat(lastMapItem.v)
                    }
                })
            }
            packagesIndex = [];
            each(requireConf.packages, function(packageConf) {
                var pkg = packageConf;
                if (typeof packageConf === "string") {
                    pkg = {
                        name: packageConf.split("/")[0],
                        location: packageConf,
                        main: "main"
                    }
                }
                pkg.location = pkg.location || pkg.name;
                pkg.main = (pkg.main || "main").replace(/\.js$/i, "");
                pkg.reg = createPrefixRegexp(pkg.name);
                packagesIndex.push(pkg)
            });
            packagesIndex.sort(descSorterByKOrName);
            urlArgsIndex = createKVSortedIndex(requireConf.urlArgs, 1);
            bundlesIndex = {};

            function bundlesIterator(id) {
                bundlesIndex[resolvePackageId(id)] = key
            }
            for (var key in requireConf.bundles) {
                each(requireConf.bundles[key], bundlesIterator)
            }
        }

        function indexRetrieve(value, index, hitBehavior) {
            each(index, function(item) {
                if (item.reg.test(value)) {
                    hitBehavior(item.v, item.k, item);
                    return false
                }
            })
        }

        function toUrl(source, baseId) {
            var extReg = /(\.[a-z0-9]+)$/i;
            var queryReg = /(\?[^#]*)$/;
            var extname = "";
            var id = source;
            var query = "";
            if (queryReg.test(source)) {
                query = RegExp.$1;
                source = source.replace(queryReg, "")
            }
            if (extReg.test(source)) {
                extname = RegExp.$1;
                id = source.replace(extReg, "")
            }
            if (baseId != null) {
                id = normalize(id, baseId)
            }
            var url = id;
            var isPathMap;
            indexRetrieve(id, pathsIndex, function(value, key) {
                url = url.replace(key, value);
                isPathMap = 1
            });
            if (!isPathMap) {
                indexRetrieve(id, packagesIndex, function(value, key, item) {
                    url = url.replace(item.name, item.location)
                })
            }
            if (!/^([a-z]{2,10}:\/)?\//i.test(url)) {
                url = requireConf.baseUrl + url
            }
            url += extname + query;
            indexRetrieve(id, urlArgsIndex, function(value) {
                url += (url.indexOf("?") > 0 ? "&" : "?") + value
            });
            return url
        }

        function createLocalRequire(baseId) {
            var requiredCache = {};

            function req(requireId, callback) {
                if (typeof requireId === "string") {
                    if (!requiredCache[requireId]) {
                        var topLevelId = normalize(requireId, baseId);
                        modTryInvokeFactory(topLevelId);
                        if (!modIs(topLevelId, MODULE_DEFINED)) {
                            throw new Error('[MODULE_MISS]"' + topLevelId + '" is not exists!')
                        }
                        requiredCache[requireId] = modModules[topLevelId].exports
                    }
                    return requiredCache[requireId]
                } else if (requireId instanceof Array) {
                    var pureModules = [];
                    var normalizedIds = [];
                    each(requireId, function(id, i) {
                        var idInfo = parseId(id);
                        var absId = normalize(idInfo.mod, baseId);
                        var resId = idInfo.res;
                        var normalizedId = absId;
                        if (resId) {
                            var trueResId = absId + "!" + resId;
                            if (resId.indexOf(".") !== 0 && bundlesIndex[trueResId]) {
                                absId = normalizedId = trueResId
                            } else {
                                normalizedId = null
                            }
                        }
                        normalizedIds[i] = normalizedId;
                        modFlagAutoDefine(absId);
                        pureModules.push(absId)
                    });
                    nativeAsyncRequire(pureModules, function() {
                        each(normalizedIds, function(id, i) {
                            if (id == null) {
                                id = normalizedIds[i] = normalize(requireId[i], baseId);
                                modFlagAutoDefine(id)
                            }
                        });
                        nativeAsyncRequire(normalizedIds, callback, baseId);
                        modAutoDefine()
                    }, baseId);
                    modAutoDefine()
                }
            }
            req.toUrl = function(id) {
                return toUrl(id, baseId || "")
            };
            return req
        }

        function normalize(id, baseId) {
            if (!id) {
                return ""
            }
            baseId = baseId || "";
            var idInfo = parseId(id);
            if (!idInfo) {
                return id
            }
            var resourceId = idInfo.res;
            var moduleId = relative2absolute(idInfo.mod, baseId);
            indexRetrieve(baseId, mappingIdIndex, function(value) {
                indexRetrieve(moduleId, value, function(mdValue, mdKey) {
                    moduleId = moduleId.replace(mdKey, mdValue)
                })
            });
            moduleId = resolvePackageId(moduleId);
            if (resourceId) {
                var mod = modIs(moduleId, MODULE_DEFINED) && actualGlobalRequire(moduleId);
                resourceId = mod && mod.normalize ? mod.normalize(resourceId, function(resId) {
                    return normalize(resId, baseId)
                }) : normalize(resourceId, baseId);
                moduleId += "!" + resourceId
            }
            return moduleId
        }

        function resolvePackageId(id) {
            each(packagesIndex, function(packageConf) {
                var name = packageConf.name;
                if (name === id) {
                    id = name + "/" + packageConf.main;
                    return false
                }
            });
            return id
        }

        function relative2absolute(id, baseId) {
            if (id.indexOf(".") !== 0) {
                return id
            }
            var segs = baseId.split("/").slice(0, -1).concat(id.split("/"));
            var res = [];
            for (var i = 0; i < segs.length; i++) {
                var seg = segs[i];
                switch (seg) {
                    case ".":
                        break;
                    case "..":
                        if (
                            res.length && res[res.length - 1] !== "..") {
                            res.pop()
                        } else {
                            res.push(seg)
                        }
                        break;
                    default:
                        seg && res.push(seg)
                }
            }
            return res.join("/")
        }

        function parseId(id) {
            var segs = id.split("!");
            if (segs[0]) {
                return {
                    mod: segs[0],
                    res: segs[1]
                }
            }
        }

        function kv2List(source, keyMatchable, allowAsterisk) {
            var list = [];
            for (var key in source) {
                if (source.hasOwnProperty(key)) {
                    var item = {
                        k: key,
                        v: source[key]
                    };
                    list.push(item);
                    if (keyMatchable) {
                        item.reg = key === "*" && allowAsterisk ? /^/ : createPrefixRegexp(key)
                    }
                }
            }
            return list
        }

        function createPrefixRegexp(
            prefix) {
            return new RegExp("^" + prefix + "(/|$)")
        }

        function each(source, iterator) {
            if (source instanceof Array) {
                for (var i = 0, len = source.length; i < len; i++) {
                    if (iterator(source[i], i) === false) {
                        break
                    }
                }
            }
        }

        function descSorterByKOrName(a, b) {
            var aValue = a.k || a.name;
            var bValue = b.k || b.name;
            if (bValue === "*") {
                return -1
            }
            if (aValue === "*") {
                return 1
            }
            return bValue.length - aValue.length
        }
        var currentlyAddingScript;
        var interactiveScript;

        function getCurrentScript() {
            if (currentlyAddingScript) {
                return currentlyAddingScript
            } else if (
                interactiveScript && interactiveScript.readyState === "interactive") {
                return interactiveScript
            }
            var scripts = document.getElementsByTagName("script");
            var scriptLen = scripts.length;
            while (scriptLen--) {
                var script = scripts[scriptLen];
                if (script.readyState === "interactive") {
                    interactiveScript = script;
                    return script
                }
            }
        }
        var headElement = document.getElementsByTagName("head")[0];
        var baseElement = document.getElementsByTagName("base")[0];
        if (baseElement) {
            headElement = baseElement.parentNode
        }

        function createScript(moduleId,
            onload) {
            var script = document.createElement("script");
            script.setAttribute("data-require-id", moduleId);
            script.src = toUrl(moduleId + ".js");
            script.async = true;
            if (script.readyState) {
                script.onreadystatechange = innerOnload
            } else {
                script.onload = innerOnload
            }

            function innerOnload() {
                var readyState = script.readyState;
                if (typeof readyState === "undefined" || /^(loaded|complete)$/.test(readyState)) {
                    script.onload = script.onreadystatechange = null;
                    script = null;
                    onload()
                }
            }
            currentlyAddingScript = script;
            baseElement ? headElement.insertBefore(script, baseElement) : headElement.appendChild(script);
            currentlyAddingScript = null
        }
        globalRequire.clearModule = function(ids) {
            if (!ids || !ids.length) {
                return
            }
            each(ids, function(id) {
                if (modModules[id]) {
                    delete modModules[id]
                }
                if (loadingModules[id]) {
                    delete loadingModules[id]
                }
                return true
            })
        };
        if (!define) {
            define = globalDefine;
            if (!require) {
                require = globalRequire
            }
            esl = globalRequire
        }
        var mainModule;
        (function() {
            var scripts = document.getElementsByTagName("script");
            var len = scripts.length;
            while (len--) {
                var script = scripts[len];
                if (mainModule = script.getAttribute("data-main")) {
                    break
                }
            }
        })();
        mainModule && setTimeout(function() {
            globalRequire([mainModule])
        }, 4)
    })(this);
    root.define = define;
    root.require = require;
    root.esl = esl
})(window);