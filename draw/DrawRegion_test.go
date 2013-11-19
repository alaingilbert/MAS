package draw

import (
  "testing"
)

//func TestDrawtile(t *testing.T) {
//  regionX, regionZ := GetRegionFromXYZ(0, 0, 0)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(1, 0, 0)
//  if regionX != 1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(2, 0, 0)
//  if regionX != 2 || regionZ != 0 { t.Fail() }
//
//  regionX, regionZ = GetRegionFromXYZ(0, 0, 1)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(1, 0, 1)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(2, 0, 1)
//  if regionX != 1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(3, 0, 1)
//  if regionX != 1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(4, 0, 1)
//  if regionX != 2 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(-1, 0, 1)
//  if regionX != -1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(-2, 0, 1)
//  if regionX != -1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(-3, 0, 1)
//  if regionX != -2 || regionZ != 0 { t.Fail() }
//
//  regionX, regionZ = GetRegionFromXYZ(0, 0, 2)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(1, 0, 2)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(2, 0, 2)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(3, 0, 2)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(4, 0, 2)
//  if regionX != 1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(8, 0, 2)
//  if regionX != 2 || regionZ != 0 { t.Fail() }
//
//  regionX, regionZ = GetRegionFromXYZ(0, 0, 3)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(7, 0, 3)
//  if regionX != 0 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(8, 0, 3)
//  if regionX != 1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(15, 0, 3)
//  if regionX != 1 || regionZ != 0 { t.Fail() }
//  regionX, regionZ = GetRegionFromXYZ(16, 0, 3)
//  if regionX != 2 || regionZ != 0 { t.Fail() }
//}


func TestStartingChunk(t *testing.T) {
  chunk := StartingChunk(-1, 0)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(0, 0)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(1, 0)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(2, 0)
  if chunk != 0 { t.Fatal(chunk) }

  chunk = StartingChunk(-2, 1)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(-1, 1)
  if chunk != 16 { t.Fatal(chunk) }
  chunk = StartingChunk(0, 1)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(1, 1)
  if chunk != 16 { t.Fatal(chunk) }
  chunk = StartingChunk(2, 1)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(3, 1)
  if chunk != 16 { t.Fatal(chunk) }

  chunk = StartingChunk(-4, 2)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(-3, 2)
  if chunk != 8 { t.Fatal(chunk) }
  chunk = StartingChunk(-2, 2)
  if chunk != 16 { t.Fatal(chunk) }
  chunk = StartingChunk(-1, 2)
  if chunk != 24 { t.Fatal(chunk) }
  chunk = StartingChunk(0, 2)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(1, 2)
  if chunk != 8 { t.Fatal(chunk) }
  chunk = StartingChunk(2, 2)
  if chunk != 16 { t.Fatal(chunk) }
  chunk = StartingChunk(3, 2)
  if chunk != 24 { t.Fatal(chunk) }

  chunk = StartingChunk(0, 3)
  if chunk != 0 { t.Fatal(chunk) }
  chunk = StartingChunk(1, 3)
  if chunk != 4 { t.Fatal(chunk) }
  chunk = StartingChunk(2, 3)
  if chunk != 8 { t.Fatal(chunk) }
  chunk = StartingChunk(3, 3)
  if chunk != 12 { t.Fatal(chunk) }
  chunk = StartingChunk(4, 3)
  if chunk != 16 { t.Fatal(chunk) }
  chunk = StartingChunk(5, 3)
  if chunk != 20 { t.Fatal(chunk) }
  chunk = StartingChunk(6, 3)
  if chunk != 24 { t.Fatal(chunk) }
  chunk = StartingChunk(7, 3)
  if chunk != 28 { t.Fatal(chunk) }
}


func TestNbChunk(t *testing.T) {
  nbChunk := NbChunk(0)
  if nbChunk != 32 { t.Fail() }
  nbChunk = NbChunk(1)
  if nbChunk != 16 { t.Fail() }
  nbChunk = NbChunk(2)
  if nbChunk != 8 { t.Fail() }
  nbChunk = NbChunk(3)
  if nbChunk != 4 { t.Fail() }
}


func TestGetScale(t *testing.T) {
  scale := GetScale(0)
  if scale != 1 { t.Fail() }
  scale = GetScale(1)
  if scale != 1 { t.Fail() }
  scale = GetScale(2)
  if scale != 2 { t.Fail() }
  scale = GetScale(3)
  if scale != 4 { t.Fail() }
  scale = GetScale(4)
  if scale != 8 { t.Fail() }
}


func TestSize(t *testing.T) {
  z := 0
  scale := GetScale(z)
  nbChunk := NbChunk(z)
  if nbChunk * 8 * scale != 256  { t.Fail() }

  z = 1
  scale = GetScale(z)
  nbChunk = NbChunk(z)
  if nbChunk * 16 * scale != 256  { t.Fail() }

  z = 2
  scale = GetScale(z)
  nbChunk = NbChunk(z)
  if nbChunk * 16 * scale != 256  { t.Fail() }

  z = 3
  scale = GetScale(z)
  nbChunk = NbChunk(z)
  if nbChunk * 16 * scale != 256  { t.Fail() }

  z = 4
  scale = GetScale(z)
  nbChunk = NbChunk(z)
  if nbChunk * 16 * scale != 256  { t.Fail() }
}


func TestBlockToSkip(t *testing.T) {
  skip := BlockToSkip(0)
  if skip != 2 { t.Fail() }

  skip = BlockToSkip(1)
  if skip != 1 { t.Fail() }

  skip = BlockToSkip(2)
  if skip != 1 { t.Fail() }

  skip = BlockToSkip(3)
  if skip != 1 { t.Fail() }
}
