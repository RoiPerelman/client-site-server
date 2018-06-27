// synchronously loads the url.
// we want to load dynamicyield before we continue loading the SPA
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

const loadDynamicYield = async (defaultSection, jsCode) => {
  console.log('loading DY with ' + defaultSection);
  console.log('loading DY with ' + jsCode);
  if (jsCode) {
    try {
      // eslint-disable-next-line
      eval(jsCode);
    } catch (e) {
      console.log(e);
    }
  }
  window.DY = window.DY || {
    scsec: defaultSection,
    API: (...args) => {
      (window.DY.API.actions = window.DY.API.actions || []).push(args);
    }
  };
  // window.DY.recommendationContext = {"type":"PRODUCT","data":["1217282-400"]};
  window.DY.recommendationContext = { type: 'HOMEPAGE', data: [] };
  await loadScript(
    `//cdn.dynamicyield.com/api/${defaultSection}/api_dynamic.js`
  );
  await loadScript(
    `//cdn.dynamicyield.com/api/${defaultSection}/api_static.js`
  );
};

export default loadDynamicYield;
