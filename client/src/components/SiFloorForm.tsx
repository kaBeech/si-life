import { useState } from "react"

// Form for creating SiFloors
export default () => {
  const [newSiFloor, setNewSiFloor] = useState("")
  const [isPending, _setIsPending] = useState(false)

  const createSiFloor = async (e: React.FormEvent) => {
    e.preventDefault();
    alert(`Creating SiFloor: ${newSiFloor}`)
  };
  return (
    <form onSubmit={createSiFloor}>
      <label>SiFloor size:</label>
      <input
        type="number"
        required
        value={newSiFloor}
        onChange={(e) => setNewSiFloor(e.target.value)}
      />
      {
        isPending ?
          <button disabled>Adding SiFloor...</button> :
          <button type="submit" >Add SiFloor</button>
      }
    </form>
  )
}
