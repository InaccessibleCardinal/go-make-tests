export function equalsIgnoreCase(text: string | undefined, other: string | undefined) {
  if (!text && !other) {
    return true;
  }
  if (!text || !other) {
    return false;
  }
  return text.localeCompare(other, undefined, { sensitivity: 'base' }) === 0;
}

export function equalsAnyIgnoreCase(text: string | undefined, ...others: (string | undefined)[]) {
  if (!text && (!others || others.some((index) => !index))) {
    return true;
  }
  if (!text) {
    return false;
  }
  return others.some((other) => equalsIgnoreCase(text, other));
}

export const includesIgnoreCase = (includedIn: string | undefined, isIncluded: string | undefined) => {
  if (includedIn === isIncluded) {
    return true;
  }
  if (!includedIn || !isIncluded) {
    return false;
  }
  return includedIn.toLowerCase().includes(isIncluded.toLowerCase());
};

export const includesAnyIgnoreCase = (
  theString: string | undefined,
  ...stringsToBeIncluded: (string | undefined)[]
) => {
  if (!theString && (!stringsToBeIncluded || stringsToBeIncluded.some((index) => !index))) {
    return true;
  }
  if (!theString) {
    return false;
  }
  return stringsToBeIncluded.some((s) => includesIgnoreCase(theString, s));
};

export const startsWithIgnoreCase = (theString: string | undefined, startsWith: string | undefined) => {
  if (!theString && !startsWith) {
    return true;
  }
  if (!theString || !startsWith) {
    return false;
  }
  return theString.toLowerCase().startsWith(startsWith.toLowerCase());
};

export const startsWithAnyIgnoreCase = (theString: string | undefined, ...startsWith: (string | undefined)[]) => {
  if (!theString && (!startsWith || startsWith.some((index) => !index))) {
    return true;
  }
  if (!theString) {
    return false;
  }
  return startsWith.some((s) => startsWithIgnoreCase(theString, s));
};
