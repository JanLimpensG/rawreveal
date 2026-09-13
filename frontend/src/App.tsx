import { useEffect, useState } from 'react'
import { OnFileDrop } from '../wailsjs/runtime'
import { ReadFile, InvertImage } from '../wailsjs/go/main/App'

function App() {

  const [base64Image, setBase64Image] = useState<string | null>(null);
  const [invertedImage, setInvertedImage] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
      OnFileDrop(( _, __, paths) => {
        console.log('Dropped files:', paths);
        setLoading(true);
        setError(null);
        ReadFile(paths[0]).then((data) => {
          console.log('File data received, length:', data.length);
          setBase64Image(data);
          setLoading(false);
        }).catch((err) => {
          console.error('Error reading file:', err);
          setError(err.message);
          setLoading(false);
        });
      }, true);
  }, [])

  const onTransformClick = () => {
    if (!base64Image) {
      setError("No image to transform");
      return;
    }
    setLoading(true);
    setError(null);
    InvertImage(base64Image).then((invertedData) => {
      console.log('Inverted image data received, length:', invertedData.length);
      setInvertedImage(invertedData);
      setLoading(false);
    }).catch((err) => {
      console.error('Error inverting image:', err);
      setError(err.message);
      setLoading(false);
    });
  }

  return (
    <>
    <div className="flex flex-col h-dvh w-dvw p-4 gap-4">
        <div className="flex flex-row h-1/2 w-full gap-4">
          <div
            className="border rounded overflow-hidden w-1/2 flex items-center justify-center bg-gray-100"
            style={{ "--wails-drop-target": "drop" } as React.CSSProperties}
          >
            {loading && <p className="text-gray-500">Loading...</p>}
            {error && <p className="text-red-500">Error: {error}</p>}
            {base64Image && !loading && (
              <img src={`data:image/png;base64,${base64Image}`} alt="Original" className="w-full h-full object-contain" />
            )}
            {!base64Image && !loading && !error && <p className="text-gray-400">Drop an image here</p>}
          </div>
          <div className="border rounded overflow-hidden w-1/2 flex items-center justify-center bg-gray-100">
            {invertedImage && !loading && (
              <img src={`data:image/png;base64,${invertedImage}`} alt="Inverted" className="w-full h-full object-contain" />
            )}
            {!invertedImage && !loading && <p className="text-gray-400">Inverted image will appear here</p>}
          </div>
        </div>
        <button className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600" onClick={onTransformClick}>
          Transform
        </button>
      </div>
    </>
  )
}

export default App
