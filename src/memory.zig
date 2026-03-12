const MAX_MEMORY_SIZE: i32 = 4096;

const Memory = struct {
    data: u8[MAX_MEMORY_SIZE],
    pub fn new() Memory {
        const self: Memory = .{};
        @memset(self.data, 0);
        return self;
    }

    pub fn write(self: *Memory, address: u12, data: u8) void {
        self.data[address] = data;
    }

    pub fn read(self: *Memory, address: u12) u8 {
        return self.data[address];
    }
};
