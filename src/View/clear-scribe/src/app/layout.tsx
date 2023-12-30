import Navbar from './components/Navbar'

import Footer from './components/Footer'

import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Clear Scribe',
  description: 'Clear Scribe is your ultimate tool for converting complex text into clear, concise, and easily understandable language.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <Navbar />
        {children}

        <Footer />

      </body>

    </html>
  )
}
