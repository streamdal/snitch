import { Handlers } from "$fresh/server.ts";
import { getConfigs } from "root/lib/fetch.ts";
import "lodash";

const cleanObject = (obj: any): any => {
  const keysToRemove = ["WasmId", "WasmFunction"];

  if (_.isArray(obj)) {
    return _.compact(obj.map(cleanObject)).filter((o) => !_.isEmpty(o));
  } else if (_.isObject(obj)) {
    return _.omitBy(
      _.mapValues(_.omit(obj, keysToRemove), cleanObject),
      (value) => _.isEmpty(value) && !_.isNumber(value),
    );
  }
  return obj;
};

export const handler: Handlers = {
  async GET() {
    const { config = {} } = await getConfigs();

    const cleanedConfig = cleanObject(config);
    return new Response(
      JSON.stringify(cleanedConfig, undefined, 2),
      {
        headers: {
          "Content-Type": "application/json",
          "Content-Disposition": "attachment; filename=streamdal-configs.json",
        },
      },
    );
  },
};
