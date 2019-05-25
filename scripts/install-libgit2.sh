cd "/var"

LG2VER="0.27.0"

wget -O libgit2-${LG2VER}.tar.gz https://github.com/libgit2/libgit2/archive/v{LG2VER}.tar.gz
tar -xzvf libgit2-${LG2VER}.tar.gz
cd libgit2-${LG2VER} && mkdir build && cd build
cmake -DTHREADSAFE=ON -DBUILD_CLAR=OFF -DCMAKE_BUILD_TYPE="RelWithDebInfo" .. && make && sudo make install
sudo ldconfig
cd "var"
rm -f libgit2-${LG2VER}.tar.gz && rm -rf libgit2-${LG2VER}