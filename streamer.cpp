#include "include/streamer.h"

extern "C" {
    int s_createDynamicObject(int modelid, const float x, const float y, const float z, const float rx, const float ry, const float rz, int worldid, int interiorid, int playerid) {
        return call<int>("streamer_createDynamicObject", modelid, x, y, z, rx, ry, rz, worldid, interiorid, playerid);
    }

    void s_destroyDynamicObject(int objectid) {
        return call<void>("streamer_destroyDynamicObject", objectid);
    }
}
