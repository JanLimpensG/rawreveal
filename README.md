# RawReveal

A lightweight desktop application for digitizing and editing film negative scans. Perfect for film photographers who want to quickly convert their scanned negatives to viewable positives with minimal hassle.

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


## Planned Features

- **TIFF and RAW File Support** - Expand beyond PNG and JPG to work with professional image formats
- **Customizable Film Stock Profiles** - Pre-configured settings optimized for different film types and emulsions
- **User Profile System** - Save and reuse personalized inversion and editing profiles for consistent results across future scans
- **Bulk Invert** - Process multiple image files in batch for efficient workflow with large scan collections
- **Image Rating & Tagging** - Organize and categorize your scanned negatives with custom tags and ratings for easy sorting and retrieval