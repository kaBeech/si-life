import { SiFloor } from "../types";

export default ({ siFloor }: { siFloor: SiFloor }) => {
  return (
    <>
      <p>Height: {siFloor.height}</p>
      <p>Width: {siFloor.width}</p>
    </>
  );
}

