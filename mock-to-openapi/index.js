/**
 * Convert item to OpenAPI schema
 * @param item
 * @returns {{format: (string), type: string, example: number}|{type: string, items: {type: ("undefined"|"object"|"boolean"|"number"|"string"|"function"|"symbol"|"bigint")}}|{format: string, type: string}|{}|{type: string, example: string}|{type: string, example: number}|{type: string, example: boolean}|{format: string, type: string, example: number}}
 */

const { Buffer } = require("node:buffer");
const { DateTime } = require("luxon");

/**
 * Detect base64 string
 * @returns {boolean}
 * @param string
 */
function isBase64(string) {
  const dec = Buffer.from(string, "base64").toString("utf8");
  return string === Buffer.from(dec, "binary").toString("base64");
}

/**
 * Detect format
 * @see https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#data-types
 * @param item
 * @returns {string|string}
 */
const getFormat = (item) => {
  // Integers range detection
  if (getType(item) === "integer") {
    if (item > -2_147_483_647 && item < 2_147_483_647) {
      return "int32";
    }

    return Number.isSafeInteger(item) ? "int64" : "unsafe";
  }

  if (getType(item) === "string") {
    if (!Number.isNaN(Date.parse(item))) {
      if (
        DateTime.fromFormat(item, "yyyy").isValid ||
        DateTime.fromFormat(item, "yyyy-MM").isValid ||
        DateTime.fromFormat(item, "yyyy-MM-dd").isValid
      ) {
        return "date";
      }

      if (
        DateTime.fromSQL(item).isValid ||
        DateTime.fromISO(item).isValid ||
        DateTime.fromHTTP(item).isValid ||
        DateTime.fromRFC2822(item).isValid
      ) {
        return "date-time";
      }
    }

    // Base64 encoded data
    if (item && isBase64(item)) {
      return "byte";
    }
  }
}
const getType = (item) => {
  if (typeof item === "string") {
    return "string";
  }

  if (typeof item === "number") {
    return Number.isInteger(item) ? "integer" : "number";
  }

  if (Object.prototype.toString.call(item) === "[object Array]") {
    return "array";
  }

  if (item && typeof item === "object") {
    return "object";
  }

  if (typeof item === "boolean") {
    return "boolean";
  }
}
const toOpenApi = (item) => {
  const oa = {};

  const format = getFormat(item);
  const type = getType(item);
  const example = item;

  switch (type) {
    case "object":
      oa.type = "object";
      oa.properties = {};
      for (const [key, value] of Object.entries(item)) {
        oa.properties[key] = toOpenApi(value);
      }

      break;
    case "array":
      return { type, items: toOpenApi(item[0]) };
    case "integer":
      return { type, format, example };
    case "number":
      return { type, example };
    case "boolean":
      return { type, example };
    case "string":
      return format ? { type, format, example } : { type, example };
    default:
      return { type: "string", format: "nullable" };
  }

  return oa;
}

module.exports = toOpenApi;
