
interface FunkyStuff {
  name: string;
  location: string;
  otherAttr: string;
}


export function processFunkyStuff(stuff: FunkyStuff): FunkyStuff {
  const newStuff = {};

  for (const key of Object.keys(stuff)) {
    newStuff[key] = stuff[key].toUpperCase();
  }
  return newStuff as FunkyStuff;
}