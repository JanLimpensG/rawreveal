# RawReveal

A lightweight desktop application for digitizing and editing film negative scans. Perfect for film photographers who want to quickly convert their scanned negatives to viewable positives with minimal hassle.

## Features

- **Instant Negative Inversion** - Convert scanned film negatives to positives with a single click
- **Histogram Analysis** - Visualize pixel frequency distribution to understand exposure and contrast
- **Drag & Drop Interface** - Simply drag your image files onto the app to load them
- **Light Image Editing** - Basic adjustments for film negative scans
- **Multi-Format Support** - Works with PNG, JPG, and other common image formats
- **Fast Performance** - Native desktop app built with Go and React for speed and reliability

## Getting Started

### Prerequisites
- Go 1.20+
- Node.js 18+

### Development

To run in live development mode with hot reload:

```bash
wails dev
```

This will launch both the frontend dev server (with Vite) and the backend Go server. Any changes to your React components or Go code will automatically refresh.

### Building for Production

To create a production-ready executable:

```bash
wails build
```

The built application will be in the `build/bin/` directory.

## Technology Stack

- **Frontend**: React 19 + TypeScript + Tailwind CSS
- **Charts**: Recharts for histogram visualization
- **Backend**: Go for fast image processing
- **Desktop Framework**: Wails for native cross-platform desktop app

## How to Use

1. **Drop an Image** - Drag and drop a scanned film negative onto the left panel
2. **View Histogram** - The histogram will automatically display the pixel frequency distribution
3. **Invert** - Click the "Transform Image" button to invert the negative to a positive
4. **Export** - Download the inverted image for further editing in your preferred photo editor

## Tips for Best Results

- Use high-quality scans for better results (at least 2400 DPI recommended)
- Ensure consistent lighting when scanning negatives
- For color negatives, the histogram shows red channel distribution
- For fine-tuned adjustments, use the histogram to guide your exposure