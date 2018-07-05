// synchronously loads the url.
// we want to load dynamicyield before we continue loading the SPA
import { history } from '../index';

const loadScript = url => {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script');
    script.type = 'text/javascript';
    script.async = false;

    script.onload = () => {
      resolve(`url script ${url} has been loaded`);
    };

    script.onerror = function() {
      reject(new Error(`Failed to load ${url}`));
    };

    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
  });
};

const getContextType = () => {
  const cutPathname = pathname => {
    if (pathname.charAt(0) === '/')
      return pathname
        .slice(1)
        .replace('Page', '')
        .toUpperCase();
  };
  return cutPathname(history.location.pathname);
};

const loadDynamicYield = async ({ section, contexts, jsCode }) => {
  window.DY = window.DY || {
    scsec: section,
    API: (...args) => {
      (window.DY.API.actions = window.DY.API.actions || []).push(args);
    }
  };
  console.log();
  window.DY.recommendationContext = {
    type: getContextType() || 'OTHER',
    data: contexts[getContextType().toLowerCase()] || []
  };
  await loadScript(`//cdn.dynamicyield.com/api/${section}/api_dynamic.js`);
  // manipulate info from api_dynamic
  if (jsCode) {
    try {
      // eslint-disable-next-line
      eval(jsCode);
    } catch (e) {
      console.log(e);
    }
  }
  // apiDynamicDYExpsManipulation();
  await loadScript(`//cdn.dynamicyield.com/api/${section}/api_static.js`);
};

// const apiDynamicDYExpsManipulation = () => {
//   // eslint-disable-next-line
//
//   // // hook manipulation
//   // DYExps.hooks.beforeSmartExecution = (tagId, tagName) => {
//   //   console.log("beforeTagExecuted", tagId, tagName)
//   // };
//   console.log(
//     'at this point we can manipulate DYExps before static script runs'
//   );
// };
export default loadDynamicYield;
