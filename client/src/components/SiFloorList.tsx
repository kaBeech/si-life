import { useState } from "react";
import { SiFloor } from "../types";
import SiFloorPreview from "./SiFloorPreview";

export default () => {
  const [isLoading, _setIsLoading] = useState(false);
  const siFloors: SiFloor[] = [
    {
      _id: 1,
      size: 100,
    },
    {
      _id: 2,
      size: 200,
    },
  ];
  return (
    <>
      <h2>SiFloors</h2>
      {isLoading ? (
        <p>Loading...</p>
      ) : (
        <ul>
          {
            siFloors.map((siFloor) => (
              <li key={siFloor.id}>
                <SiFloorPreview siFloor={siFloor} />
              </li>
            ))
          }
        </ul>
      )}
    </>
  );
}
