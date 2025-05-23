type hmap struct {
    count     int           // number of key-value pairs
    flags     uint8
    B         uint8         // log₂(number of buckets)
    noverflow uint16        // approximate number of overflow buckets
    hash0     uint32        // hash seed

    buckets    unsafe.Pointer // array of buckets
    oldbuckets unsafe.Pointer // for incremental resize
    nevacuate  uintptr        // progress counter for evacuating oldbuckets
    extra *mapextra // rarely used stuff
}

type bmap struct {
    tophash [8]uint8      // Top 8 bits of hash for quick match
    keys    [8]keyType    // Actual keys
    values  [8]valueType  // Actual values
    overflow *bmap        // Pointer to overflow bucket if full
}

HashFunction(key,seed)