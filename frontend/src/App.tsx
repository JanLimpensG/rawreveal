import { useEffect, useState } from 'react'
import { OnFileDrop } from '../wailsjs/runtime'
import { ReadFile } from '../wailsjs/go/main/App'

function App() {

  const [base64Image, setBase64Image] = useState<string | null>(null);

  useEffect(() => {
      OnFileDrop(( _, __, paths) => {
        console.log('Dropped files:', paths);
        ReadFile(paths[0]).then((data) => {
          console.log('File data:', data);
          setBase64Image(data);
        });
      }, true);
  }, [])

  return (
    <>
    <div className="flex flex-col h-dvh w-dvw p-4 gap-4">
        <div className="flex flex-row h-1/2 w-full gap-4">
          <div
            className="border rounded overflow-hidden w-1/2"
            style={{ "--wails-drop-target": "drop" } as React.CSSProperties}
          >
            { base64Image &&
              <img src={`data:image/png;base64,${base64Image}`} alt="Dropped" className="w-full h-full object-contain" />
            } 
          </div>
          <div className="border rounded overflow-hidden w-1/2">
          </div>
        </div>
        <button className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"> Tranform </button>
      </div>
    </>
  )
}

export default App
