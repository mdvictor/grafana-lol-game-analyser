export const splitCamelCase = (camelCaseStr: string) => {
  const splitStr = camelCaseStr.replace(/([a-z0-9])([A-Z])/g, '$1 $2');

  return splitStr.charAt(0).toUpperCase() + splitStr.slice(1);
};
