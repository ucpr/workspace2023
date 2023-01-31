const pino = require("pino");

const labels = pino().levels.labels;
const logger = pino({
  level: "info",
  mixin: (_, level) => {
    return {
      severity: labels[level].toUpperCase(),
    };
  },
  mixinMergeStrategy(mergeObject, mixinObject) {
    return Object.assign({ jsonPayload: mergeObject }, mixinObject);
  },
});


logger.info({ tag: 1 }, 'error massage')             // 重要度: INFO    | ErrorReport: 
logger.info(new Error('error message'))  // 重要度: INFO    | ErrorReport:
logger.warn('error massage')             // 重要度: WARNING | ErrorReport: 
logger.warn(new Error('error message'))  // 重要度: WARNING | ErrorReport:
logger.error('error massage')            // 重要度: ERROR   | ErrorReport: 
logger.error(new Error('error message')) // 重要度: ERROR   | ErrorReport: ✅ 

