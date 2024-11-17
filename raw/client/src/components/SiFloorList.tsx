import { BASE_URL } from "../App";
import { SiFloor } from "../types";
import SiFloorPreview from "./SiFloorPreview";
import { useQuery } from "@tanstack/react-query";

export default () => {
  const { data: siFloor, isLoading } = useQuery<SiFloor>({
    queryKey: ["siFloors"],
    queryFn: async () => {
      try {
        const res = await fetch(BASE_URL + "/sifloor/1");
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error);
        }
        return data || [];
      } catch (error) {
        console.error(error);
      }
    },
  });
  return (
    <>
      <h2>SiFloors</h2>
      {isLoading ? (
        <p>Loading...</p>
      ) : (
        <ul>
          {
            <li >
              <SiFloorPreview siFloor={siFloor!} />
            </li>
          }
        </ul>
      )}
    </>
  );
}
