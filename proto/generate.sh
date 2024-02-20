cd proto
buf generate
cd ..

cp -r github.com/wfblockchain/noble-cctp/* ./
rm -rf github.com
