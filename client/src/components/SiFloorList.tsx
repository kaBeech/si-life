import { SiFloor } from "../types";
import SiFloorPreview from "./SiFloorPreview";
import { useQuery } from "@tanstack/react-query";

export default () => {
  const { data: siFloors, isLoading } = useQuery<SiFloor[]>({
    queryKey: ["siFloors"],
    queryFn: async () => {
      try {
        const res = await fetch("http://localhost:4000/api/siFloors");
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
            siFloors!.map((siFloor) => (
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
