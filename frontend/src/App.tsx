import { useEffect, useState } from 'react'
import { OnFileDrop } from '../wailsjs/runtime'
import { LoadImage  } from '../wailsjs/go/main/App'
//import { AreaChart, Area, Legend, ResponsiveContainer } from 'recharts'

function App() {

  const [base64Image, setBase64Image] = useState<string | null>(null);
  //const [histogramData, setHistogramData] = useState<{ r: number[]; g: number[]; b: number[]; gray: number[] } | null>(null);
  //const [invertedImage, setInvertedImage] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
 // const [visibleAreas, setVisibleAreas] = useState({ red: true, green: true, blue: true, gray: true });

  useEffect(() => {
      OnFileDrop(( _, __, paths) => {
        console.log('Dropped files:', paths);
        setLoading(true);
        setError(null);
        LoadImage(paths[0]).then((data: any) => {
          setBase64Image(data);
          // Store histogram with r, g, b, gray channels
          //setHistogramData(data.histogram || null);
          setLoading(false);
        }).catch((err) => {
          console.error('Error reading file:', err);
          setError(err.message);
          setLoading(false);
        });
      }, true);
  }, [])

  /*
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

  // Create chart data from histogram
  const chartData = histogramData
    ? Array.from({ length: 64 }, (_, i) => {
        const idx = i * 4;
        return {
          intensity: idx,
          red: histogramData.r[idx] || 0,
          green: histogramData.g[idx] || 0,
          blue: histogramData.b[idx] || 0,
          gray: histogramData.gray[idx] || 0,
        };
      })
    : null;

  function handleLegendClicked(data: any): void {
    const dataKey = data.dataKey;
    setVisibleAreas(prev => ({
      ...prev,
      [dataKey]: !prev[dataKey as keyof typeof visibleAreas]
    }));
  }
*/
  return (
    <div className="min-h-screen">
        <div className="grid grid-cols-1  gap-6 mb-8">
          {/* Image Drop Zone */}
          <div className="grid-cols-2 bg-white rounded-2xl ">
            <h2 className="text-xl font-semibold text-gray-800 mb-4">Original Image</h2>
            <div
              className="border-3 border-dashed border-blue-300 rounded-xl overflow-hidden flex items-center justify-center bg-gradient-to-br from-blue-50 to-blue-100 min-h-96"
              style={{ "--wails-drop-target": "drop" } as React.CSSProperties}
            >
              {loading && (
                <div className="flex flex-col items-center gap-2">
                  <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
                  <p className="text-gray-600">Processing...</p>
                </div>
              )}
              {error && <p className="text-red-500 font-semibold text-center px-4">{error}</p>}
              {base64Image && !loading && (
                <img src={`data:image/png;base64,${base64Image}`} alt="Original" className="max-w-full max-h-96 object-contain" />
              )}
              {!base64Image && !loading && !error && (
                <div className="text-center">
                  <p className="text-gray-500 text-lg">Drop an image here</p>
                </div>
              )}
            </div>
          </div>
        </div>

        {/* Inverted Image and Chart 
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
          {/* Inverted Image 
          <div className="bg-white rounded-2xl shadow-lg p-6 hover:shadow-xl transition">
            <h2 className="text-xl font-semibold text-gray-800 mb-4">Inverted Image</h2>
            <div className="border-3 border-dashed border-purple-300 rounded-xl overflow-hidden flex items-center justify-center bg-gradient-to-br from-purple-50 to-purple-100 min-h-96">
              {invertedImage && !loading && (
                <img src={`data:image/png;base64,${invertedImage}`} alt="Inverted" className="max-w-full max-h-96 object-contain" />
              )}
              {!invertedImage && !loading && (
                <div className="text-center">
                  <p className="text-gray-500 text-lg">Inverted image appears here</p>
                  <p className="text-gray-400 text-sm mt-2">Click Transform to invert</p>
                </div>
              )}
            </div>
          </div>

          {/* Chart 
          <div className="bg-white rounded-2xl shadow-lg p-6 hover:shadow-xl transition">
            <h2 className="text-xl font-semibold text-gray-800 mb-4">Histogram</h2>
            {chartData && chartData.length > 0 ? (
              <ResponsiveContainer width="100%" height={300}>
                <AreaChart
                  data={chartData}
                  margin={{ top: 20, right: 30, left: 0, bottom: 20 }}
                  
                >
                  <Legend onClick={handleLegendClicked}/>
                  <Area 
                    type="monotone" 
                    dataKey="red" 
                    fill="#ef4444" 
                    stroke="#dc2626"
                    fillOpacity={visibleAreas.red ? 0.6 : 0}
                    strokeOpacity={visibleAreas.red ? 1 : 0}
                    isAnimationActive={false}
                  />
                  <Area 
                    type="monotone" 
                    dataKey="green" 
                    fill="#22c55e" 
                    stroke="#16a34a"
                    fillOpacity={visibleAreas.green ? 0.6 : 0}
                    strokeOpacity={visibleAreas.green ? 1 : 0}
                    isAnimationActive={false}
                  />
                  <Area 
                    type="monotone" 
                    dataKey="blue" 
                    fill="#3b82f6" 
                    stroke="#1d4ed8"
                    fillOpacity={visibleAreas.blue ? 0.6 : 0}
                    strokeOpacity={visibleAreas.blue ? 1 : 0}
                    isAnimationActive={false}
                  />
                  <Area 
                    type="monotone" 
                    dataKey="gray" 
                    fill="#818181" 
                    stroke="#444444"
                    fillOpacity={visibleAreas.gray ? 0.3 : 0}
                    strokeOpacity={visibleAreas.gray ? 1 : 0}
                    isAnimationActive={false}
                  />
                </AreaChart>
              </ResponsiveContainer>
            ) : (
              <div className="h-80 flex items-center justify-center">
                <p className="text-gray-500">Upload an image to see histogram</p>
              </div>
            )}
          </div>
        </div>

         Transform Button 
        <div>
          <button
            onClick={onTransformClick}
            disabled={!base64Image || loading}
            className="w-full py-3 px-6 bg-gradient-to-r from-blue-500 to-indigo-600 text-white font-bold text-lg rounded-xl hover:shadow-lg hover:from-blue-600 hover:to-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition transform hover:scale-105 active:scale-95"
          >
            {loading ? "Processing..." : "Transform Image"}
          </button>
        </div>
        */}
    </div>
  )
}

export default App
